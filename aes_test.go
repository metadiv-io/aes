package aes_test

import (
	"github.com/metadiv-io/aes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("test: aes", func() {

	var encryptKey = "McQfTjWnZr4u7x!A%D*G-KaNdRgUkXp2"

	It("should be able to encrypt and decrypt (bytes)", func() {
		encrypted := aes.EncryptToBytes("hello", encryptKey)
		decrypted := aes.DecryptFromBytes(encrypted, encryptKey)
		Expect(decrypted).To(Equal("hello"))
	})

	It("should be able to encrypt and decrypt (base64)", func() {
		encrypted := aes.EncryptToBase64("hello", encryptKey)
		decrypted := aes.DecryptFromBase64(encrypted, encryptKey)
		Expect(decrypted).To(Equal("hello"))
	})
})
