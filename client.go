package tg

import (
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/LukeEuler/trx-go/config"
	"github.com/LukeEuler/trx-go/trx/api"
)

/*
https://developers.tron.network/docs/networks
*/

var (
	once           sync.Once
	wallet         api.WalletClient
	walletSolidity api.WalletSolidityClient
)

func initClient() {
	f := func() {
		conf := config.Get()
		conn, err := grpc.Dial(conf.Net.Wallet, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		wallet = api.NewWalletClient(conn)

		conn, err = grpc.Dial(conf.Net.WalletSolidity, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		walletSolidity = api.NewWalletSolidityClient(conn)
	}
	once.Do(f)
}
