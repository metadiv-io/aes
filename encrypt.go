package aes

import (
	"crypto/aes"
	"encoding/base64"
)

var Encrypt = new(encrypt)

type encrypt struct{}

// BytesToBytes encrypts the given byte slice using the provided key and returns the encrypted byte slice.
func (e *encrypt) BytesToBytes(src []byte, key string) []byte {
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

// BytesToString encrypts the given byte slice using the provided key and returns the encrypted string.
func (e *encrypt) BytesToString(src []byte, key string) string {
	return string(e.BytesToBytes(src, key))
}

// BytesToBase64 encrypts the given byte slice using the provided key and returns the encrypted base64 encoded string.
func (e *encrypt) BytesToBase64(src []byte, key string) string {
	return base64.StdEncoding.EncodeToString(e.BytesToBytes(src, key))
}

// StringToBytes encrypts the given string using the provided key and returns the encrypted byte slice.
func (e *encrypt) StringToBytes(str, key string) []byte {
	src := []byte(str)
	return e.BytesToBytes(src, key)
}

// StringToString encrypts the given string using the provided key and returns the encrypted string.
func (e *encrypt) StringToString(str, key string) string {
	return string(e.StringToBytes(str, key))
}

// StringToBase64 encrypts the given string using the provided key and returns the encrypted base64 encoded string.
func (e *encrypt) StringToBase64(str, key string) string {
	return base64.StdEncoding.EncodeToString(e.StringToBytes(str, key))
}
