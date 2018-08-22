package consensus

import (
	"bytes"
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/thetatoken/ukulele/blockchain"
	"github.com/thetatoken/ukulele/p2p"
)

var _ Engine = &DefaultEngine{}

// DefaultEngine is the default implementation of the Engine interface.
type DefaultEngine struct {
	chain   *blockchain.Chain
	network p2p.Network

	incoming        chan interface{}
	finalizedBlocks chan *blockchain.Block

	// TODO: persist state
	// Consensus state
	highestCCBlock     *blockchain.ExtendedBlock
	lastFinalizedBlock *blockchain.ExtendedBlock
	tip                *blockchain.ExtendedBlock
	lastVoteHeight     uint32
	voteLog            map[uint32]blockchain.Vote     // level -> vote
	collectedVotes     map[string]*blockchain.VoteSet // block hash -> votes
	epochManager       *EpochManager
	epoch              uint32
	validatorManager   ValidatorManager

	// Strategies
	proposerStrategy ProposerStrategy
	replicaStrategy  ReplicaStrategy
}

// NewEngine creates a instance of DefaultEngine.
func NewEngine(chain *blockchain.Chain, network p2p.Network, validators *ValidatorSet) *DefaultEngine {
	e := &DefaultEngine{
		chain:   chain,
		network: network,

		incoming:        make(chan interface{}, 5000),
		finalizedBlocks: make(chan *blockchain.Block, 5000),

		highestCCBlock:     chain.Root,
		lastFinalizedBlock: chain.Root,
		tip:                chain.Root,
		voteLog:            make(map[uint32]blockchain.Vote),
		collectedVotes:     make(map[string]*blockchain.VoteSet),
		validatorManager:   NewRotatingValidatorManager(validators),
		epochManager:       NewEpochManager(),
		epoch:              0,

		proposerStrategy: &DefaultProposerStrategy{},
		replicaStrategy:  &DefaultReplicaStrategy{},
	}
	e.proposerStrategy.Init(e)
	e.replicaStrategy.Init(e)
	e.epochManager.Init(e)
	network.AddMessageHandler(e)
	return e
}

// ID returns the identifier of current node.
func (e *DefaultEngine) ID() string {
	return e.network.ID()
}

// Chain return a pointer to the underlying chain store.
func (e *DefaultEngine) Chain() *blockchain.Chain {
	return e.chain
}

// Network returns a pointer to the underlying network.
func (e *DefaultEngine) Network() p2p.Network {
	return e.network
}

// SetProposerStrategy allows us to change proposerStrategy.
func (e *DefaultEngine) SetProposerStrategy(s ProposerStrategy) {
	s.Init(e)
	e.proposerStrategy = s
}

// SetReplicaStrategy allows us to change replicaStrategy.
func (e *DefaultEngine) SetReplicaStrategy(s ReplicaStrategy) {
	s.Init(e)
	e.replicaStrategy = s
}

// Start is the main event loop.
func (e *DefaultEngine) Start(ctx context.Context) {
	go e.epochManager.Start(ctx)
	go e.mainLoop(ctx)
}

func (e *DefaultEngine) mainLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-e.incoming:
			switch m := msg.(type) {
			case Proposal:
				e.handleProposal(m)
			case blockchain.Vote:
				e.handleVote(m)
			default:
				log.Errorf("Unknown message type: %v", m)
			}
		case newEpoch := <-e.epochManager.C:
			e.enterNewEpoch(newEpoch)
		}
	}
}

func (e *DefaultEngine) enterNewEpoch(newEpoch uint32) {
	e.epoch = newEpoch
	e.proposerStrategy.EnterNewEpoch(newEpoch)
	e.replicaStrategy.EnterNewEpoch(newEpoch)
}

// HandleMessage implements p2p.MessageHandler interface.
func (e *DefaultEngine) HandleMessage(network p2p.Network, msg interface{}) {
	e.incoming <- msg
}

func (e *DefaultEngine) handleProposal(proposal Proposal) {
	e.replicaStrategy.HandleProposal(proposal)
}

func (e *DefaultEngine) handleVote(vote blockchain.Vote) {
	e.proposerStrategy.HandleVote(vote)
}

// setTip sets the block to extended from by next proposal. Currently we use the highest block among highestCCBlock's
// descendants as the fork-choice rule.
func (e *DefaultEngine) setTip() *blockchain.ExtendedBlock {
	ret, _ := e.highestCCBlock.FindDeepestDescendant()
	e.tip = ret
	return ret
}

// getTip return the block to be extended from.
func (e *DefaultEngine) getTip() *blockchain.ExtendedBlock {
	return e.tip
}

// FinalizedBlocks returns a channel that will be published with finalized blocks by the engine.
func (e *DefaultEngine) FinalizedBlocks() chan *blockchain.Block {
	return e.finalizedBlocks
}

func (e *DefaultEngine) processCCBlock(ccBlock *blockchain.ExtendedBlock) {
	log.WithFields(log.Fields{"id": e.ID(), "ccBlock": ccBlock, "c.epoch": e.epoch}).Debug("Start processing ccBlock")
	defer log.WithFields(log.Fields{"id": e.ID(), "ccBlock": ccBlock, "c.epoch": e.epoch}).Debug("Done processing ccBlock")

	if ccBlock.Height > e.highestCCBlock.Height {
		log.WithFields(log.Fields{"id": e.ID(), "ccBlock": ccBlock}).Debug("Updating highestCCBlock since ccBlock.Height > e.highestCCBlock.Height")
		e.highestCCBlock = ccBlock
	}

	if ccBlock.Parent.CommitCertificate != nil {
		e.finalizeBlock(ccBlock.Parent)
	}

	if ccBlock.Epoch >= e.epoch {
		log.WithFields(log.Fields{"id": e.ID(), "ccBlock": ccBlock, "e.epoch": e.epoch}).Debug("Advancing epoch")
		newEpoch := ccBlock.Epoch + 1
		e.enterNewEpoch(newEpoch)
		e.epochManager.SetEpoch(newEpoch)
	}
}

func (e *DefaultEngine) finalizeBlock(block *blockchain.ExtendedBlock) {
	// Skip blocks that have already published.
	if bytes.Compare(block.Hash, e.lastFinalizedBlock.Hash) == 0 {
		return
	}

	log.WithFields(log.Fields{"id": e.ID(), "block.Hash": block.Hash}).Info("Finalizing block")
	defer log.WithFields(log.Fields{"id": e.ID(), "block.Hash": block.Hash}).Info("Done Finalized block")

	e.lastFinalizedBlock = block

	select {
	case e.finalizedBlocks <- block.Block:
	default:
	}
}