package tg

import (
	"fmt"

	"github.com/LukeEuler/trx-go/config"
	"github.com/LukeEuler/trx-go/key"
)

func NewKeys() {
	conf := config.Get()
	if !conf.Keys.Enable || conf.Keys.Number <= 0 {
		return
	}
	for i := 0; i < conf.Keys.Number; i++ {
		k := key.NewKey()
		fmt.Printf("%s %s\n", k.PrivateKey(), k.Address())
	}
}
