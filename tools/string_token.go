package tools

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/gtank/cryptopasta"
)

func StringToken(input string, password string, saltHex string) string {
	// 這裡的實作與 FileToken 類似，只是將檔案讀取改為直接使用輸入的字串
	sum := sha256.Sum256([]byte(input))
	key := DeriveKey(password, saltHex)

	ciphertext, err := cryptopasta.Encrypt(sum[:], key)
	if err != nil {
		panic("failed to encrypt: " + err.Error())
	}

	encoded := hex.EncodeToString(ciphertext)
	return encoded
}
