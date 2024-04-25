package aes

/*
DecryptStringToBytes decrypts the encrypted bytes with the given key and returns the decrypted bytes.
*/
func DecryptStringToBytes(encrypted string, key string) []byte {
	src := []byte(encrypted)
	return DecryptBytesToBytes(src, key)
}

/*
DecryptStringToString decrypts the encrypted bytes with the given key and returns the decrypted text.
*/
func DecryptStringToString(encrypted string, key string) string {
	return string(DecryptStringToBytes(encrypted, key))
}

/*
DecryptBase64ToString decrypts the encrypted base64 string with the given key and returns the decrypted text.
*/
func DecryptBase64ToString(encrypted string, key string) string {
	return string(DecryptBase64ToBytes(encrypted, key))
}
