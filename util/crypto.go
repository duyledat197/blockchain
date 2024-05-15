package util

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewPrivateKey generates a new private key.
//
// No parameters.
// Returns *ecdsa.PrivateKey, string, error.
func NewPrivateKey() (*ecdsa.PrivateKey, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, "", status.Errorf(codes.Internal, "unable to generate private key: %v", err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	priKeyString := hexutil.Encode(privateKeyBytes)[2:]

	return privateKey, priKeyString, nil
}

// PubKeyFromPrivKey generates a public key from a given private key.
//
// privateKey: The private key from which to generate the public key.
// Returns a string representing the public key and a pointer to the ecdsa.PublicKey.
func PubKeyFromPrivKey(privateKey *ecdsa.PrivateKey) (string, *ecdsa.PublicKey) {
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	return hexutil.Encode(publicKeyBytes)[4:], publicKeyECDSA
}
