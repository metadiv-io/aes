package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// EncryptToBytes encrypts the text with the given key and returns the encrypted bytes.
func EncryptToBytes(text, key string) []byte {
	src := []byte(text)
	return EncryptBytesToBytes(src, key)
}

// EncryptBytesToBytes encrypts the bytes with the given key and returns the encrypted bytes.
func EncryptBytesToBytes(src []byte, key string) []byte {
	cipher := getCipher(key)
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))

	for bs, be := 0, cipher.BlockSize(); bs <= len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return encrypted
}

// DecryptFromBytes decrypts the encrypted bytes with the given key and returns the decrypted text.
func DecryptFromBytes(encrypted []byte, key string) string {
	return string(DecryptBytesFromBytes(encrypted, key))
}

// DecryptBytesFromBytes decrypts the encrypted bytes with the given key and returns the decrypted bytes.
func DecryptBytesFromBytes(encrypted []byte, key string) []byte {
	cipher := getCipher(key)

	decrypted := make([]byte, len(encrypted))
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

// EncryptToBase64 encrypts the text with the given key and returns the encrypted base64 string.
func EncryptToBase64(text, key string) string {
	return base64.StdEncoding.EncodeToString(EncryptToBytes(text, key))
}

// DecryptFromBase64 decrypts the encrypted base64 string with the given key and returns the decrypted text.
func DecryptFromBase64(encrypted string, key string) string {
	decrypted, _ := base64.StdEncoding.DecodeString(encrypted)
	return DecryptFromBytes(decrypted, key)
}

func getCipher(key string) cipher.Block {
	if key == "" {
		panic("aes key is empty") // must not happen
	}
	cipher, _ := aes.NewCipher(generateKey([]byte(key)))
	return cipher
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
