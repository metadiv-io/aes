package aes

import "encoding/base64"

/*
EncryptStringToBytes encrypts the text with the given key and returns the encrypted bytes.
*/
func EncryptStringToBytes(str, key string) []byte {
	src := []byte(str)
	return EncryptBytesToBytes(src, key)
}

/*
EncryptStringToString encrypts the text with the given key and returns the encrypted text.
*/
func EncryptStringToString(str, key string) string {
	return string(EncryptStringToBytes(str, key))
}

/*
EncryptStringToBase64 encrypts the text with the given key and returns the encrypted base64 string.
*/
func EncryptStringToBase64(str, key string) string {
	return base64.StdEncoding.EncodeToString(EncryptStringToBytes(str, key))
}
