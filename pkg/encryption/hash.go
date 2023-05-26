package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func EncryptData(data []byte) (string, error) {
	key := []byte("0NSThAOHUa2HetD6MyBOKNED6537ofjN")

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	// Encode ciphertext ke base64 URL-safe
	encryptedData := base64.RawURLEncoding.EncodeToString(ciphertext)
	return encryptedData, nil
}
