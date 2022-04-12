package key

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"math/big"

	"github.com/pkg/errors"
	"github.com/shengdoushi/base58"

	btcSecp256k1 "github.com/LukeEuler/trx-go/key/btc_secp256k1"
)

type KeyPair struct {
	private []byte
	address string
}

func NewKey() *KeyPair {
	bs := make([]byte, 32)
	_, _ = io.ReadFull(rand.Reader, bs)
	return &KeyPair{
		private: bs,
	}
}

func NewKeyFromBytes(bs []byte) (*KeyPair, error) {
	if len(bs) != 32 {
		return nil, errors.New("invalid bytes")
	}
	return &KeyPair{
		private: bs,
	}, nil
}

func NewKeyFromHex(raw string) (*KeyPair, error) {
	bs, err := hex.DecodeString(raw)
	if err != nil {
		return nil, err
	}
	return NewKeyFromBytes(bs)
}

func (k *KeyPair) PrivateKey() string {
	return hex.EncodeToString(k.private)
}

func (k *KeyPair) Address() string {
	if len(k.address) != 0 {
		return k.address
	}
	private, _ := btcSecp256k1.PrivKeyFromBytes(btcSecp256k1.S256(), k.private)
	privKey := private.ToECDSA()

	address := PubkeyToEthAddressBytes(privKey.PublicKey)

	addressTron := make([]byte, 0)
	addressTron = append(addressTron, TronBytePrefix)
	addressTron = append(addressTron, address...)

	if addressTron[0] == 0 {

		k.address = new(big.Int).SetBytes(addressTron).String()
	} else {
		k.address = EncodeCheck(addressTron)
	}
	return k.address
}

const (
	TronBytePrefix = byte(0x41)
)

func EncodeCheck(input []byte) string {
	h256h0 := sha256.New()
	h256h0.Write(input)
	h0 := h256h0.Sum(nil)

	h256h1 := sha256.New()
	h256h1.Write(h0)
	h1 := h256h1.Sum(nil)

	inputCheck := input
	inputCheck = append(inputCheck, h1[:4]...)

	return Encode(inputCheck)
}

func Encode(input []byte) string {
	return base58.Encode(input, base58.BitcoinAlphabet)
}
