package main

import (
	"flag"

	"github.com/LukeEuler/dolly/log"

	tg "github.com/LukeEuler/trx-go"
	"github.com/LukeEuler/trx-go/config"
)

func main() {
	log.AddConsoleOut(5)

	configFile := flag.String("c", "config.toml", "set the config file path")
	flag.Parse()
	log.Entry.Infof("config file: %s", *configFile)
	config.New(*configFile)

	tg.Do()

	tg.NewKeys()
	tg.GetBalance()
	tg.Transfer()
}
