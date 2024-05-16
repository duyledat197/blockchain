package eth_util

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"reflect"
	"regexp"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/sha3"

	"openmyth/blockchain/pkg/eth_client"
)

// PublicKeyBytesToAddress converts a public key bytes slice to an Ethereum address.
//
// Parameters:
// - publicKey: The public key bytes slice to convert.
// Return type: common.Address.
func PublicKeyBytesToAddress(publicKey []byte) common.Address {
	var buf []byte

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKey[1:]) // remove EC prefix 04
	buf = hash.Sum(nil)
	address := buf[12:]

	return common.HexToAddress(hex.EncodeToString(address))
}

// IsAddressContract checks if the given address corresponds to a smart contract on the Ethereum blockchain.
//
// Parameters:
// - client: The Ethereum client used to interact with the blockchain.
// - ctx: The context for the function.
// - addr: The address to be checked.
// Return type: bool indicating if the address is a contract, error if any error occurs during the check.
func IsAddressContract(ctx context.Context, client eth_client.IClient, addr string) (bool, error) {
	address := common.HexToAddress(addr)
	bytecode, err := client.CodeAt(ctx, address, nil) // nil is latest block
	if err != nil {
		return false, err
	}

	return len(bytecode) > 0, nil
}

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()

	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// CalcGasCost calculate gas cost given gas limit (units) and gas price (wei)
func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBig := big.NewInt(int64(gasLimit))
	return gasLimitBig.Mul(gasLimitBig, gasPrice)
}

// SigRSV signatures R S V returned as arrays
func SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)

	return R, S, V
}

func VerifySignature(hexPrivateKey, signatureStr, nonce string) bool {

	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	if err != nil {
		return false
	}

	// keccak256 hash of the data
	dataBytes := []byte(nonce)
	hashData := crypto.Keccak256Hash(dataBytes)
	signature, err := hexutil.Decode(signatureStr[:len(signatureStr)-2] + "00")
	if err != nil {
		return false
	}
	sigPublicKey, err := crypto.Ecrecover(hashData.Bytes(), signature)
	if err != nil {
		return false
	}

	publicKey := privateKey.Public()
	publicKeyBytes := crypto.FromECDSAPub(publicKey.(*ecdsa.PublicKey))

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	if !matches {
		return false
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(hashData.Bytes(), signature)
	if err != nil {
		return false
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hashData.Bytes(), signatureNoRecoverID)

	return verified
}
