package aes

import "encoding/base64"

/*
DecryptBytesToBytes decrypts the encrypted bytes with the given key and returns the decrypted bytes.
*/
func DecryptBytesToBytes(encrypted []byte, key string) []byte {
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

/*
DecryptBytesToString decrypts the encrypted bytes with the given key and returns the decrypted string.
*/
func DecryptBytesToString(encrypted []byte, key string) string {
	return string(DecryptBytesToBytes(encrypted, key))
}

/*
DecryptBase64ToBytes decrypts the encrypted base64 string with the given key and returns the decrypted bytes.
*/
func DecryptBase64ToBytes(encrypted string, key string) []byte {
	decrypted, _ := base64.StdEncoding.DecodeString(encrypted)
	return DecryptBytesToBytes(decrypted, key)
}
