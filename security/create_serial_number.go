package security

import (
	"crypto/rand"
	"math/big"
)

// CreateSerialNumber creates a random serial number that is a maximum length of
// 20 bytes. See https://tools.ietf.org/html/rfc3280#appendix-B
func CreateSerialNumber() (*big.Int, error) {
	limit := new(big.Int).Lsh(big.NewInt(1), 128)
	return rand.Int(rand.Reader, limit)
}
