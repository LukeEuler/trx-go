package tg

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/LukeEuler/dolly/log"

	"github.com/LukeEuler/trx-go/common"
	"github.com/LukeEuler/trx-go/config"
	"github.com/LukeEuler/trx-go/trx/core"
)

func GetBalance() {
	conf := config.Get()
	if !conf.Balance.Enable {
		return
	}

	for _, address := range conf.Balance.Address {
		getBalance(address)
		fmt.Printf("%s/#/address/%s\n\n", conf.Net.Show, address)
	}
}

func getBalance(address string) {
	conf := config.Get()
	if !conf.Balance.Enable {
		return
	}

	initClient()

	var err error
	account := new(core.Account)
	account.Address, err = common.DecodeCheck(address)
	if err != nil {
		panic(err)
	}

	acc, err := walletSolidity.GetAccount(context.Background(), account)
	if err != nil {
		panic(err)
	}

	if !bytes.Equal(acc.Address, account.Address) {
		log.Entry.Warn("account not found")
		return
	}
	fmt.Printf("%s balance:\n", address)
	fmt.Printf("  %d TRX\n", acc.GetBalance())
	if conf.Balance.Trc10.Enable {
		if conf.Balance.Trc10.ShowAll {
			for k, v := range acc.AssetV2 {
				fmt.Printf("  %s: %d\n", k, v)
			}
		} else {
			for _, k := range conf.Balance.Trc10.IDs {
				v, ok := acc.AssetV2[k]
				if !ok {
					fmt.Printf("  %s: not fount\n", k)
				} else {
					fmt.Printf("  %s: %d\n", k, v)
				}
			}
		}
	}

	if !conf.Balance.Trc20.Enable {
		return
	}

	for _, item := range conf.Balance.Trc20.Assets {
		list := strings.SplitN(item, " ", 2)
		asset := list[0]
		name := "$"
		if len(list) > 1 {
			name = list[1]
		}
		account.Address, err = common.DecodeCheck(address)
		if err != nil {
			log.Fatal(err)
		}
		addrHex := hex.EncodeToString(account.Address)

		data := trc20BalanceOf + "0000000000000000000000000000000000000000000000000000000000000000"[len(addrHex)-2:] + addrHex[2:]

		// fromBsxx, _ := hex.DecodeString("410000000000000000000000000000000000000000")
		contractBs, err := common.DecodeCheck(asset)
		if err != nil {
			log.Fatal(err)
		}

		dataBs, err := hex.DecodeString(data)
		if err != nil {
			log.Entry.Fatal(err)
		}

		ct := &core.TriggerSmartContract{
			OwnerAddress:    account.Address, // 这里可以是 hex.DecodeString("410000000000000000000000000000000000000000")
			ContractAddress: contractBs,
			Data:            dataBs,
		}

		result, err := walletSolidity.TriggerConstantContract(context.Background(), ct)
		if err != nil {
			log.Entry.Fatal(err)
		}

		if result.Result.Code > 0 {
			log.Entry.Error(string(result.Result.GetMessage()))
		}

		amountBytes := result.GetConstantResult()[0]
		amount := big.NewInt(0).SetBytes(amountBytes)
		fmt.Printf("  %s: %s %s\n", asset, amount, name)
	}
}
