package tg

import (
	"context"

	"github.com/LukeEuler/dolly/log"

	"github.com/LukeEuler/trx-go/trx/api"
)

func GetBlockNow() {
	initClient()
	b, err := walletSolidity.GetNowBlock2(context.Background(), &api.EmptyMessage{})
	if err != nil {
		log.Entry.Fatal(err)
	}
	log.Entry.Infof("height: %d", b.GetBlockHeader().GetRawData().GetNumber())
}
