package tg

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/LukeEuler/dolly/log"
	"google.golang.org/protobuf/proto"

	"github.com/LukeEuler/trx-go/common"
	"github.com/LukeEuler/trx-go/config"
	"github.com/LukeEuler/trx-go/trx"
	"github.com/LukeEuler/trx-go/trx/api"
	"github.com/LukeEuler/trx-go/trx/core"
)

func Transfer() {
	conf := config.Get()

	unSupportTxTypes := make([]string, 0, 3)
	for txType := range conf.Transfer {
		switch txType {
		default:
			unSupportTxTypes = append(unSupportTxTypes, txType)
		case "trx", "trc10", "trc20":
		}
	}
	if len(unSupportTxTypes) != 0 {
		log.Entry.Warnf("invalid transfer types: %v", unSupportTxTypes)
	}

	trxTfs := conf.Transfer["trx"]
	for _, item := range trxTfs {
		transferTrx(item)
	}

	trc10Tfs := conf.Transfer["trc10"]
	for _, item := range trc10Tfs {
		transferTrc10(item)
	}

	trc20Tfs := conf.Transfer["trc20"]
	for _, item := range trc20Tfs {
		transferTrc20(item)
	}
}

func transferTrx(tf config.Transfr) {
	if !tf.Enable {
		return
	}
	conf := config.Get()
	key, ok := conf.KeyPair[tf.From]
	if !ok {
		log.Entry.Warnf("can not found key for: %s", tf.From)
	}
	initClient()

	amount, ok := big.NewInt(0).SetString(tf.Amount, 10)
	if !ok {
		log.Entry.Errorf("invalid amount: [%s]", tf.Amount)
	}

	contract := &core.TransferContract{}

	var err error
	contract.OwnerAddress, err = common.DecodeCheck(tf.From)
	if err != nil {
		log.Fatal(err)
	}
	contract.ToAddress, err = common.DecodeCheck(tf.To)
	if err != nil {
		log.Fatal(err)
	}
	contract.Amount = amount.Int64()

	tx, err := wallet.CreateTransaction2(context.Background(), contract)
	if err != nil {
		log.Entry.Fatal(err)
	}

	signAndBroadcast(tx, key)
}

func transferTrc10(tf config.Transfr) {
	if !tf.Enable {
		return
	}
	conf := config.Get()
	key, ok := conf.KeyPair[tf.From]
	if !ok {
		log.Entry.Warnf("can not found key for: %s", tf.From)
	}
	initClient()

	amount, ok := big.NewInt(0).SetString(tf.Amount, 10)
	if !ok {
		log.Entry.Errorf("invalid amount: [%s]", tf.Amount)
	}

	var err error
	contract := &core.TransferAssetContract{}

	contract.OwnerAddress, err = common.DecodeCheck(tf.From)
	if err != nil {
		log.Fatal(err)
	}

	contract.ToAddress, err = common.DecodeCheck(tf.To)
	if err != nil {
		log.Fatal(err)
	}

	contract.AssetName = []byte(tf.Asset)
	contract.Amount = amount.Int64()

	tx, err := wallet.TransferAsset2(context.Background(), contract)
	if err != nil {
		log.Entry.Fatal(err)
	}

	signAndBroadcast(tx, key)
}

func transferTrc20(tf config.Transfr) {
	if !tf.Enable {
		return
	}
	conf := config.Get()
	key, ok := conf.KeyPair[tf.From]
	if !ok {
		log.Entry.Warnf("can not found key for: %s", tf.From)
	}
	initClient()

	amount, ok := big.NewInt(0).SetString(tf.Amount, 10)
	if !ok {
		log.Entry.Errorf("invalid amount: [%s]", tf.Amount)
	}

	fromBytes, err := common.DecodeCheck(tf.From)
	if err != nil {
		log.Fatal(err)
	}
	toBytes, err := common.DecodeCheck(tf.To)
	if err != nil {
		log.Fatal(err)
	}
	contractBytes, err := common.DecodeCheck(tf.Asset)
	if err != nil {
		log.Fatal(err)
	}
	toHex := hex.EncodeToString(toBytes)

	ab := common.LeftPadBytes(amount.Bytes(), 32)

	data := trc20TransferMethodSignature + "0000000000000000000000000000000000000000000000000000000000000000"[len(toHex)-2:] + toHex[2:]
	data += hex.EncodeToString(ab)
	dataBytes, err := hex.DecodeString(data)
	if err != nil {
		log.Entry.Fatal(err)
	}

	ct := &core.TriggerSmartContract{
		OwnerAddress:    fromBytes,
		ContractAddress: contractBytes,
		Data:            dataBytes,
	}

	tx, err := wallet.TriggerContract(context.Background(), ct)
	if err != nil {
		log.Entry.Fatal(err)
	}

	tx.Transaction.RawData.FeeLimit = tf.FeeLimit
	rawData, err := proto.Marshal(tx.Transaction.GetRawData())
	if err != nil {
		log.Entry.Fatal(err)
	}

	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	tx.Txid = hash

	signAndBroadcast(tx, key)
}

func signAndBroadcast(tx *api.TransactionExtention, key string) {
	if proto.Size(tx) == 0 {
		log.Entry.Fatal("bad transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		log.Entry.Fatal(string(tx.GetResult().GetMessage()))
	}

	signTx, err := trx.SignTransaction(tx.Transaction, key)
	if err != nil {
		log.Fatal(err)
	}

	result, err := wallet.BroadcastTransaction(context.Background(), signTx)
	if err != nil {
		log.Entry.Fatal(err)
	}
	if !result.GetResult() {
		log.Entry.Fatal(string(result.GetMessage()))
	}
	if result.GetCode() != api.Return_SUCCESS {
		log.Entry.Fatalf("%d %s", result.Code, string(result.GetMessage()))
	}

	hash := hex.EncodeToString(tx.GetTxid())

	fmt.Println("tx hash: " + hash)
	conf := config.Get()
	fmt.Printf("url: %s/#/transaction/%s\n", conf.Net.Show, hash)
}
