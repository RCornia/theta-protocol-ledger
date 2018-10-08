package types

import "github.com/thetatoken/ukulele/common"

// FIXME

// func TestPubkey(t *testing.T) {
// 	assert := assert.New(t)

// 	_, pubkey1, err := crypto.GenerateKeyPair()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Test conversion to/from proto message.
// 	pk := PublicKeyToProto(pubkey1)
// 	pubkey2 := PublicKeyFromProto(pk)
// 	assert.EqualValues(pubkey1, pubkey2)

// 	// Test conversion to/from bytes.
// 	b, err := ToBytes(pubkey1)
// 	assert.Nil(err)
// 	var pubkey3 crypto.PublicKey
// 	err = FromBytes(b, &pubkey3)
// 	assert.Nil(err)
// 	assert.EqualValues(*pubkey1, pubkey3)

// 	// Verify bytes are deterministic.
// 	b2, err := ToBytes(pubkey1)
// 	assert.Nil(err)
// 	assert.EqualValues(b, b2)
// }

// func TestPrivkey(t *testing.T) {
// 	assert := assert.New(t)

// 	privKey, _, err := crypto.GenerateKeyPair()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Test conversion to/from proto message.
// 	pk := PrivateKeyToProto(privKey)
// 	privkey2 := PrivateKeyFromProto(pk)
// 	assert.EqualValues(privKey, privkey2)

// 	// Test conversion to/from bytes.
// 	b, err := ToBytes(privKey)
// 	assert.Nil(err)
// 	var privkey3 crypto.PrivateKey
// 	err = FromBytes(b, &privkey3)
// 	assert.Nil(err)
// 	assert.EqualValues(*privKey, privkey3)

// 	// Verify bytes are deterministic.
// 	b2, err := ToBytes(privKey)
// 	assert.Nil(err)
// 	assert.EqualValues(b, b2)
// }

// func TestSignature(t *testing.T) {
// 	assert := assert.New(t)

// 	var b [64]byte
// 	for i := 0; i < len(b); i++ {
// 		b[i] = byte(i)
// 	}

// 	privKey, _, err := crypto.GenerateKeyPair()
// 	if err != nil {
// 		panic(err)
// 	}

// 	sig1, err := privKey.Sign(b[:])
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Test conversion to/from proto message.
// 	msg := SignatureToProto(sig1)
// 	sig2 := SignatureFromProto(msg)
// 	assert.EqualValues(sig1, sig2)

// 	// Test conversion to/from bytes.
// 	bb, err := ToBytes(sig1)
// 	assert.Nil(err)
// 	var sig3 crypto.Signature
// 	err = FromBytes(bb, &sig3)
// 	assert.Nil(err)
// 	assert.EqualValues(sig1, &sig3)

// 	// Verify bytes are deterministic.
// 	bb2, err := ToBytes(sig1)
// 	assert.Nil(err)
// 	assert.EqualValues(bb, bb2)
// }

// func TestOverSpendingProof(t *testing.T) {
// 	assert := assert.New(t)

// 	subjs := []OverspendingProof{
// 		{},
// 		{
// 			ReserveSequence: 1,
// 			ServicePayments: []ServicePaymentTx{{
// 				Fee:             Coin{Denom: "ThetaWei", Amount: 123},
// 				Gas:             123,
// 				Source:          TxInput{Address: getTestAddress("123")},
// 				Target:          TxInput{Address: getTestAddress("456")},
// 				PaymentSequence: 1,
// 				ReserveSequence: 1,
// 			}},
// 		},
// 	}

// 	for _, subj := range subjs {
// 		// Test conversion to/from bytes.
// 		b, err := ToBytes(&subj)
// 		assert.Nil(err)
// 		subj2 := &OverspendingProof{}
// 		err = FromBytes(b, subj2)
// 		assert.Nil(err)
// 		assert.EqualValues(&subj, subj2)

// 		// Verify bytes are deterministic.
// 		b2, err := ToBytes(&subj)
// 		assert.Nil(err)
// 		assert.EqualValues(b, b2)
// 	}
// }

// func TestCoinsSerialization(t *testing.T) {
// 	assert := assert.New(t)

// 	cc := []Coins{Coins(nil), {}, {{Denom: "ThetaWei", Amount: 123}}}

// 	for _, c := range cc {
// 		// Test conversion to/from proto message.
// 		msg := CoinsToProto(c)
// 		c2 := CoinsFromProto(msg)
// 		assert.EqualValues(c, c2)
// 	}
// }

// func TestAccount(t *testing.T) {
// 	assert := assert.New(t)

// 	_, pubkey1, err := crypto.GenerateKeyPair()
// 	if err != nil {
// 		panic(err)
// 	}

// 	rf := ReservedFund{
// 		// TargetAddresses: data.Bytes("target_address"),
// 		ReserveSequence: 123,
// 		EndBlockHeight:  456,
// 		InitialFund:     Coins{{Denom: "ThetaWei", Amount: 789}},
// 	}
// 	accounts := []Account{{
// 		PubKey:                 pubkey1,
// 		Sequence:               31,
// 		LastUpdatedBlockHeight: 22,
// 		Balance:                Coins{{Denom: "ThetaWei", Amount: 123}},
// 		ReservedFunds:          []ReservedFund{rf},
// 	}, {
// 	// Test with empty fields.
// 	}}

// 	for _, account1 := range accounts {
// 		// Test conversion to/from proto message.
// 		msg := AccountToProto(&account1)

// 		account2 := AccountFromProto(msg)
// 		assert.EqualValues(&account1, account2)

// 		// Test conversion to/from bytes.
// 		b, err := ToBytes(&account1)
// 		assert.Nil(err)
// 		account3 := &Account{}
// 		err = FromBytes(b, account3)
// 		assert.Nil(err)
// 		assert.EqualValues(&account1, account3)

// 		// Verify bytes are deterministic.
// 		b2, err := ToBytes(&account1)
// 		assert.Nil(err)
// 		assert.EqualValues(b, b2)
// 	}
// }

// func TestInput(t *testing.T) {
// 	assert := assert.New(t)

// 	sk, pk, err := crypto.GenerateKeyPair()
// 	if err != nil {
// 		panic(err)
// 	}

// 	var b [64]byte
// 	for i := 0; i < len(b); i++ {
// 		b[i] = byte(i)
// 	}
// 	sig, err := sk.Sign(b[:])
// 	if err != nil {
// 		panic(err)
// 	}

// 	inputs := []TxInput{{
// 		Sequence: 123,
// 	}, {
// 		Address:   getTestAddress("123"),
// 		Coins:     Coins{{Denom: "ThetaWei", Amount: 456}},
// 		PubKey:    pk,
// 		Signature: sig,
// 	}}

// 	for _, input1 := range inputs {
// 		// Test conversion to/from proto message.
// 		msg := InputToProto(&input1)
// 		input2 := InputFromProto(msg)
// 		assert.EqualValues(&input1, input2)
// 	}
// }

// func TestOutput(t *testing.T) {
// 	assert := assert.New(t)

// 	outputs := []TxOutput{{}, {
// 		Address: getTestAddress("123"),
// 		Coins:   Coins{{Denom: "ThetaWei", Amount: 456}},
// 	}}

// 	for _, output1 := range outputs {
// 		// Test conversion to/from proto message.
// 		msg := OutputToProto(&output1)
// 		output2 := OutputFromProto(msg)
// 		assert.EqualValues(&output1, output2)
// 	}
// }

// func TestSplit(t *testing.T) {
// 	assert := assert.New(t)

// 	addr, err := hex.DecodeString("D7D25858609A250BCD698CBAA3DB6B285586657C")
// 	assert.Equal(err, nil)
// 	var address common.Address
// 	copy(address[:], addr)

// 	split1 := Split{
// 		Address:    address,
// 		Percentage: 40,
// 	}

// 	msg := SplitToProto(&split1)
// 	split2 := SplitFromProto(msg)
// 	assert.EqualValues(&split1, split2)
// }

// func TestSplitContract(t *testing.T) {
// 	assert := assert.New(t)

// 	addr, err := hex.DecodeString("D7D25858609A250BCD698CBAA3DB6B285586657C")
// 	assert.Equal(err, nil)
// 	var address common.Address
// 	copy(address[:], addr)

// 	split := Split{
// 		Address:    address,
// 		Percentage: 40,
// 	}

// 	splitContract1 := SplitContract{
// 		ResourceID:     []byte("rid0000001"),
// 		Splits:         []Split{split},
// 		EndBlockHeight: 1006,
// 	}

// 	msg := SplitContractToProto(&splitContract1)
// 	splitContract2 := SplitContractFromProto(msg)
// 	assert.EqualValues(&splitContract1, splitContract2)
// }

// func TestTx(t *testing.T) {
// 	assert := assert.New(t)

// 	txs := []Tx{
// 		&CoinbaseTx{},
// 		&CoinbaseTx{
// 			Proposer:    TxInput{Address: getTestAddress("123")},
// 			Outputs:     []TxOutput{{Address: getTestAddress("456")}, {Address: getTestAddress("888")}, {Address: getTestAddress("999")}},
// 			BlockHeight: uint64(999),
// 		},

// 		&SlashTx{},
// 		&SlashTx{
// 			Proposer:        TxInput{Address: getTestAddress("123")},
// 			SlashedAddress:  getTestAddress("456"),
// 			SlashProof:      []byte("789"),
// 			ReserveSequence: 1,
// 		},

// 		&SendTx{},
// 		&SendTx{
// 			Fee:     Coin{Denom: "ThetaWei", Amount: 123},
// 			Gas:     123,
// 			Inputs:  []TxInput{{Address: getTestAddress("123")}, {Address: getTestAddress("798")}},
// 			Outputs: []TxOutput{{Address: getTestAddress("456")}, {Address: getTestAddress("888")}, {Address: getTestAddress("999")}},
// 		},

// 		&ReserveFundTx{},
// 		&ReserveFundTx{
// 			Fee:         Coin{Denom: "ThetaWei", Amount: 123},
// 			Gas:         123,
// 			Source:      TxInput{Address: getTestAddress("123")},
// 			Collateral:  Coins{{Denom: "ThetaWei", Amount: 456}},
// 			ResourceIDs: [][]byte{[]byte("789")},
// 			Duration:    1,
// 		},

// 		&ReleaseFundTx{},
// 		&ReleaseFundTx{
// 			Fee:             Coin{Denom: "ThetaWei", Amount: 123},
// 			Gas:             123,
// 			Source:          TxInput{Address: getTestAddress("123")},
// 			ReserveSequence: 1,
// 		},

// 		&ServicePaymentTx{},
// 		&ServicePaymentTx{
// 			Fee:             Coin{Denom: "ThetaWei", Amount: 123},
// 			Gas:             123,
// 			Source:          TxInput{Address: getTestAddress("123")},
// 			Target:          TxInput{Address: getTestAddress("456")},
// 			PaymentSequence: 1,
// 			ReserveSequence: 2,
// 		},

// 		&SplitContractTx{},
// 		&SplitContractTx{
// 			Fee:        Coin{Denom: "ThetaWei", Amount: 123},
// 			Gas:        123,
// 			ResourceID: []byte("rid789"),
// 			Initiator:  TxInput{Address: getTestAddress("123")},
// 			Splits:     []Split{Split{Address: getTestAddress("456"), Percentage: 40}, Split{Address: getTestAddress("777"), Percentage: 20}},
// 			Duration:   1000,
// 		},

// 		//&UpdateValidatorsTx{},
// 	}

// 	for _, tx := range txs {
// 		// Test conversion to/from bytes.
// 		b := TxToBytes(tx)
// 		tx2, err := TxFromBytes(b)
// 		assert.Nil(err)
// 		assert.EqualValues(tx, tx2)

// 		// Verify bytes are deterministic.
// 		b2 := TxToBytes(tx)
// 		assert.EqualValues(b, b2)
// 	}

// 	// Special test case for UpdateValidatosTx
// 	_, pubkey1, err := crypto.GenerateKeyPair()
// 	if err != nil {
// 		panic(err)
// 	}

// 	vaStake := uint64(1)
// 	va := core.NewValidator(pubkey1.ToBytes(), vaStake)
// 	tx := &UpdateValidatorsTx{
// 		Proposer:   TxInput{Address: getTestAddress("123")},
// 		Validators: []*core.Validator{&va},
// 	}
// 	// Test conversion to/from bytes.
// 	b := TxToBytes(tx)
// 	tx2_, err := TxFromBytes(b)
// 	tx2 := tx2_.(*UpdateValidatorsTx)
// 	assert.Nil(err)
// 	assert.EqualValues(tx.Proposer, tx2.Proposer)
// 	assert.Equal(len(tx.Validators), len(tx2.Validators))
// 	assert.EqualValues(*tx.Validators[0], *tx2.Validators[0])

// 	// Verify bytes are deterministic.
// 	b2 := TxToBytes(tx)
// 	assert.EqualValues(b, b2)
// }

func getTestAddress(addr string) common.Address {
	var address common.Address
	copy(address[:], addr)
	return address
}
