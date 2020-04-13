package whitecoin_txsigner

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"math/big"
	"sync"
)

type curveParam struct {
	elliptic.CurveParams
}

var (
	initonce sync.Once
	three    = new(big.Int).SetUint64(3)
)

var (
	ErrPrivateKeyIllegal = errors.New("Invalid private key data!")
	ErrUnknownCurve      = errors.New("Unknown curve type!")
	ErrMessageIllegal    = errors.New("Invalid message data!")
	ErrPublicKeyIllegal  = errors.New("Invalid public key data!")
)

func Signer(privateKey, hash []byte) ([]byte, byte, error) {

	if privateKey == nil || len(privateKey) != 32 {
		return nil, 0, ErrPrivateKeyIllegal
	}

	if hash == nil || len(hash) != 32 {
		return nil, 0, ErrMessageIllegal
	}

	var k1curve *secp256k1Curve
	privateKeyBig := new(big.Int).SetBytes(privateKey)
	priv := new(ecdsa.PrivateKey)

	if privateKeyBig.Cmp(big.NewInt(0)) == 0 {
		return nil, 0, ErrPrivateKeyIllegal
	}
	priv.D = privateKeyBig
	k1curve = secp256k1

	if privateKeyBig.Cmp(k1curve.Params().N) >= 0 {
		return nil, 0, ErrPrivateKeyIllegal
	}

	priv.PublicKey.Curve = k1curve

	priv.PublicKey.X, priv.PublicKey.Y = k1curve.ScalarBaseMult(privateKey)

	signature := make([]byte, 64)

	r, s, v, err := signecdsa(rand.Reader, priv, hash)
	if err != nil {
		return nil, 0, err
	}

	rBytes := r.Bytes()
	sBytes := s.Bytes()

	copy(signature[32-len(rBytes):32], rBytes)
	copy(signature[64-len(sBytes):64], sBytes)

	return signature, v, nil
}
