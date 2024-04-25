package aes

import (
	"crypto/aes"
	"encoding/base64"
)

/*
EncryptBytesToBytes encrypts the src bytes with the given key and returns the encrypted bytes.
*/
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

/*
EncryptBytesToString encrypts the src bytes with the given key and returns the encrypted string.
*/
func EncryptBytesToString(src []byte, key string) string {
	return string(EncryptBytesToBytes(src, key))
}

/*
EncryptBytesToBase64 encrypts the src bytes with the given key and returns the encrypted base64 string.
*/
func EncryptBytesToBase64(src []byte, key string) string {
	return base64.StdEncoding.EncodeToString(EncryptBytesToBytes(src, key))
}
