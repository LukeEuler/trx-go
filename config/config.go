package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/LukeEuler/dolly/log"
)

var conf = new(config)

type config struct {
	Net struct {
		Wallet         string `toml:"wallet"`
		WalletSolidity string `toml:"wallet_solidity"`
		Show           string `toml:"show"`
	} `toml:"net"`
	Keys struct {
		Enable bool `toml:"enable"`
		Number int  `toml:"number"`
	} `toml:"keys"`
	Balance struct {
		Enable  bool     `toml:"enable"`
		Address []string `toml:"address"`
		Trc10   struct {
			Enable  bool     `toml:"enable"`
			ShowAll bool     `toml:"show_all"`
			IDs     []string `toml:"ids"`
		} `toml:"trc10"`
		Trc20 struct {
			Enable bool     `toml:"enable"`
			Assets []string `toml:"assets"`
		} `toml:"trc20"`
	} `toml:"balance"`
	KeyPair  map[string]string    `toml:"key_pair"`
	Transfer map[string][]Transfr `toml:"transfer"`
}

type Transfr struct {
	Enable   bool   `toml:"enable"`
	From     string `toml:"from"`
	To       string `toml:"to"`
	Asset    string `toml:"asset"`
	Amount   string `toml:"amount"`
	FeeLimit int64  `toml:"fee_limit"`
}

func New(configPath string) {
	bs, err := os.ReadFile(configPath)
	if err != nil {
		log.Entry.Fatal(err)
	}
	_, err = toml.Decode(string(bs), conf)
	if err != nil {
		log.Entry.Fatal(err)
	}
}

func Get() *config {
	return conf
}
