package trx

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"math/big"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	btcSecp256k1 "github.com/LukeEuler/trx-go/key/btc_secp256k1"
	ethSecp256k1 "github.com/LukeEuler/trx-go/key/eth_secp256k1"
	"github.com/LukeEuler/trx-go/trx/core"
)

func decodePrivKeyByHex(privateKey string) (*ecdsa.PrivateKey, error) {
	bs, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	private, _ := btcSecp256k1.PrivKeyFromBytes(btcSecp256k1.S256(), bs)
	privKey := private.ToECDSA()
	return privKey, nil
}

func SignTransaction(transaction *core.Transaction, privateKey string) (*core.Transaction, error) {
	priv, err := decodePrivKeyByHex(privateKey)
	if err != nil {
		return nil, err
	}
	defer zeroKey(priv)
	rawData, err := proto.Marshal(transaction.GetRawData())
	if err != nil {
		return nil, errors.Wrap(err, "proto marshal tx raw data error")
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	signature, err := Sign(hash, priv)
	if err != nil {
		return nil, errors.Wrap(err, "sign error")
	}
	transaction.Signature = append(transaction.Signature, signature)
	return transaction, nil
}

// zeroKey zeroes a private key in memory.
func zeroKey(k *ecdsa.PrivateKey) {
	b := k.D.Bits()
	for i := range b {
		b[i] = 0
	}
}

// https://github.com/ethereum/go-ethereum/blob/master/crypto/signature_cgo.go

// DigestLength sets the signature digest exact length
const DigestLength = 32

const (
	// number of bits in a big.Word
	wordBits = 32 << (uint64(^big.Word(0)) >> 63)
	// number of bytes in a big.Word
	wordBytes = wordBits / 8
)

// Sign calculates an ECDSA signature.
//
// This function is susceptible to chosen plaintext attacks that can leak
// information about the private key that is used for signing. Callers must
// be aware that the given digest cannot be chosen by an adversery. Common
// solution is to hash any input before calculating the signature.
//
// The produced signature is in the [R || S || V] format where V is 0 or 1.
func Sign(digestHash []byte, prv *ecdsa.PrivateKey) (sig []byte, err error) {
	if len(digestHash) != DigestLength {
		return nil, errors.Errorf("hash is required to be exactly %d bytes (%d)", DigestLength, len(digestHash))
	}
	seckey := PaddedBigBytes(prv.D, prv.Params().BitSize/8)
	defer zeroBytes(seckey)
	return ethSecp256k1.Sign(digestHash, seckey)
}

// PaddedBigBytes encodes a big integer as a big-endian byte slice. The length
// of the slice is at least n bytes.
func PaddedBigBytes(bigint *big.Int, n int) []byte {
	if bigint.BitLen()/8 >= n {
		return bigint.Bytes()
	}
	ret := make([]byte, n)
	ReadBits(bigint, ret)
	return ret
}

// ReadBits encodes the absolute value of bigint as big-endian bytes. Callers must ensure
// that buf has enough space. If buf is too short the result will be incomplete.
func ReadBits(bigint *big.Int, buf []byte) {
	i := len(buf)
	for _, d := range bigint.Bits() {
		for j := 0; j < wordBytes && i > 0; j++ {
			i--
			buf[i] = byte(d)
			d >>= 8
		}
	}
}

func zeroBytes(bytes []byte) {
	for i := range bytes {
		bytes[i] = 0
	}
}
