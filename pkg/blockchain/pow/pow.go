package pow

import (
	"log/slog"
	"math"
	"math/big"

	"openmyth/blockchain/pkg/blockchain/block"
	"openmyth/blockchain/util"
)

const (
	targetBits = 5 // difficult value
	maxNonce   = math.MaxInt64
)

var target *big.Int

// getTarget returns the target difficulty for Proof of Work calculations.
func getTarget() *big.Int {
	if target != nil {
		return target
	}

	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	return target
}

// ProofOfWork represents a Proof of Work object.
//
// It contains the block that the proof of work is being calculated for, as well as
// the target difficulty for the proof of work.
type ProofOfWork struct {
	block  *block.Block
	target *big.Int
}

// NewProofOfWork creates a new ProofOfWork object for the given block.
func NewProofOfWork(b *block.Block) *ProofOfWork {
	return &ProofOfWork{
		block:  b,
		target: getTarget(),
	}
}

// Calculate calculates the proof of work for the given block.
//
// It iterates through nonce values until a hash value is found that meets the target difficulty.
// The calculated nonce and hash value are returned.
func (pow *ProofOfWork) Calculate() (int64, []byte) {
	nonce := int64(0)

	slog.Debug("Mining the block containing ", slog.Any("data", pow.block.Data))

	for nonce < maxNonce {
		if pow.isNonceValid(nonce) {
			break
		}

		nonce++
	}

	return nonce, pow.prepareHashData(nonce)
}

// ValidateBlock checks if the proof of work for the given block is valid.
//
// It calculates the hash of the block's data and nonce, compares it to the target difficulty,
// and returns true if the hash is less than the target, indicating a valid proof of work.
//
// Returns:
// - bool: True if the proof of work is valid, false otherwise.
func (pow *ProofOfWork) ValidateBlock() bool {
	return pow.isNonceValid(pow.block.Nonce)
}

// prepareHashData prepares the hash data for the proof of work calculation.
//
// It takes the nonce value as input and returns a byte slice containing the hash data.
// The hash data is calculated by concatenating the previous block hash, block data, block timestamp, target bits, and nonce.
// The hash data is then passed to the util.HashSHA256 function to calculate the SHA256 hash.
func (pow *ProofOfWork) prepareHashData(nonce int64) []byte {
	return util.HashSHA256(
		pow.block.PrevBlockHash,
		pow.block.Data,
		util.IntToHex(pow.block.Timestamp),
		util.IntToHex(int64(targetBits)),
		util.IntToHex(nonce),
	)
}

// isNonceValid checks if the given nonce produces a hash value less than the target difficulty.
func (pow *ProofOfWork) isNonceValid(nonce int64) bool {
	var hashInt big.Int

	hash := pow.prepareHashData(nonce)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}
