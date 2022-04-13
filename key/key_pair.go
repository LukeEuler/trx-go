package key

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"

	"github.com/pkg/errors"
	"github.com/shengdoushi/base58"

	"github.com/LukeEuler/trx-go/common"
	btcSecp256k1 "github.com/LukeEuler/trx-go/key/btc_secp256k1"
)

type KeyPair struct {
	priv *ecdsa.PrivateKey
}

func NewKey() (*KeyPair, error) {
	key, err := ecdsa.GenerateKey(btcSecp256k1.S256(), rand.Reader)
	return &KeyPair{priv: key}, errors.WithStack(err)
}

func NewKeyFromBytes(bs []byte) (*KeyPair, error) {
	if len(bs) != 32 {
		return nil, errors.New("invalid bytes")
	}
	curve := btcSecp256k1.S256()
	x, y := curve.ScalarBaseMult(bs)
	priv := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     x,
			Y:     y,
		},
		D: new(big.Int).SetBytes(bs),
	}
	return &KeyPair{priv: priv}, nil
}

func NewKeyFromHex(hexKey string) (*KeyPair, error) {
	bs, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, err
	}
	return NewKeyFromBytes(bs)
}

func (k *KeyPair) PrivateKey() string {
	return hex.EncodeToString(common.LeftPadBytes(k.priv.D.Bytes(), 32))
}

func (k *KeyPair) Address() string {
	bs := PubkeyToEthAddressBytes(k.priv.PublicKey)

	addressTron := make([]byte, 0)
	addressTron = append(addressTron, TronBytePrefix)
	addressTron = append(addressTron, bs...)

	var address string
	if addressTron[0] == 0 {
		address = new(big.Int).SetBytes(addressTron).String()
	} else {
		address = EncodeCheck(addressTron)
	}
	return address
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
