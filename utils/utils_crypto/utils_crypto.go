package utils_crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func CryptoSha256(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	return mdStr
}
