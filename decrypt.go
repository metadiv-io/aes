package aes

import "encoding/base64"

var Decrypt = new(decrypt)

type decrypt struct{}

// BytesToBytes decrypts the given byte slice using the provided key and returns the decrypted byte slice.
func (d *decrypt) BytesToBytes(src []byte, key string) []byte {
	cipher := getCipher(key)

	decrypted := make([]byte, len(src))
	for bs, be := 0, cipher.BlockSize(); bs < len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], src[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

// BytesToString decrypts the given byte slice using the provided key and returns the decrypted string.
func (d *decrypt) BytesToString(src []byte, key string) string {
	return string(d.BytesToBytes(src, key))
}

// StringToBytes decrypts the given encrypted string using the provided key and returns the decrypted byte slice.
func (d *decrypt) StringToBytes(encrypted string, key string) []byte {
	src := []byte(encrypted)
	return d.BytesToBytes(src, key)
}

// StringToString decrypts the given encrypted string using the provided key and returns the decrypted string.
func (d *decrypt) StringToString(encrypted string, key string) string {
	return string(d.StringToBytes(encrypted, key))
}

// Base64ToBytes decrypts the given base64 encoded string using the provided key and returns the decrypted byte slice.
func (d *decrypt) Base64ToBytes(src string, key string) []byte {
	decrypted, _ := base64.StdEncoding.DecodeString(src)
	return d.BytesToBytes(decrypted, key)
}

// Base64ToString decrypts the given base64 encoded string using the provided key and returns the decrypted string.
func (d *decrypt) Base64ToString(encrypted string, key string) string {
	return string(d.Base64ToBytes(encrypted, key))
}
