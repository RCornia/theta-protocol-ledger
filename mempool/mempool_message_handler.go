package mempool

import (
	"fmt"

	"github.com/thetatoken/ukulele/serialization/rlp"

	"github.com/thetatoken/ukulele/common"
	"github.com/thetatoken/ukulele/p2p/types"

	dp "github.com/thetatoken/ukulele/dispatcher"
)

//
// MempoolMessageHandler handles the messages received over the
// ChannelIDTransaction channel
//
type MempoolMessageHandler struct {
	mempool *Mempool
}

// CreateMempoolMessageHandler create an instance of the MempoolMessageHandler
func CreateMempoolMessageHandler(mempool *Mempool) *MempoolMessageHandler {
	return &MempoolMessageHandler{
		mempool: mempool,
	}
}

// GetChannelIDs implements the p2p.MessageHandler interface
func (mmp *MempoolMessageHandler) GetChannelIDs() []common.ChannelIDEnum {
	return []common.ChannelIDEnum{
		common.ChannelIDTransaction,
	}
}

// ParseMessage implements the p2p.MessageHandler interface
func (mmp *MempoolMessageHandler) ParseMessage(peerID string, channelID common.ChannelIDEnum, rawMessageBytes common.Bytes) (types.Message, error) {
	var dataResponse dp.DataResponse
	rlp.DecodeBytes(rawMessageBytes, &dataResponse)

	// TODO: verify the checksum
	mptx := mempoolTransaction{
		rawTransaction: dataResponse.Payload,
	}
	message := types.Message{
		PeerID:    peerID,
		ChannelID: channelID,
		Content:   mptx,
	}
	return message, nil
}

// HandleMessage implements the p2p.MessageHandler interface
func (mmp *MempoolMessageHandler) HandleMessage(message types.Message) error {
	if message.ChannelID != common.ChannelIDTransaction {
		return fmt.Errorf("Invalid channel for MempoolMessageHandler: %v", message.ChannelID)
	}
	mptx := message.Content.(*mempoolTransaction)
	err := mmp.mempool.ProcessTransaction(mptx)
	return err
}