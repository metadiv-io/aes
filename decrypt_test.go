package aes

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Decrypt", func() {
	const (
		testKey     = "0123456789abcdef" // 16-byte key for AES-128
		testString  = "Hello, World!"
		emptyString = ""
	)

	var (
		encrypted    []byte
		encryptedStr string
		base64Str    string
	)

	BeforeEach(func() {
		// Prepare encrypted data for tests
		encrypted = Encrypt.BytesToBytes([]byte(testString), testKey)
		encryptedStr = Encrypt.StringToString(testString, testKey)
		base64Str = Encrypt.StringToBase64(testString, testKey)
	})

	Describe("BytesToBytes", func() {
		It("should decrypt bytes correctly", func() {
			decrypted := Decrypt.BytesToBytes(encrypted, testKey)
			Expect(string(decrypted)).To(Equal(testString))
		})

		It("should handle empty input", func() {
			emptyEncrypted := Encrypt.BytesToBytes([]byte{}, testKey)
			decrypted := Decrypt.BytesToBytes(emptyEncrypted, testKey)
			Expect(decrypted).To(BeEmpty())
		})
	})

	Describe("BytesToString", func() {
		It("should decrypt bytes to string correctly", func() {
			decrypted := Decrypt.BytesToString(encrypted, testKey)
			Expect(decrypted).To(Equal(testString))
		})

		It("should handle empty input", func() {
			emptyEncrypted := Encrypt.BytesToBytes([]byte{}, testKey)
			decrypted := Decrypt.BytesToString(emptyEncrypted, testKey)
			Expect(decrypted).To(BeEmpty())
		})
	})

	Describe("StringToBytes", func() {
		It("should decrypt string to bytes correctly", func() {
			decrypted := Decrypt.StringToBytes(encryptedStr, testKey)
			Expect(string(decrypted)).To(Equal(testString))
		})

		It("should handle empty input", func() {
			emptyEncrypted := Encrypt.StringToString(emptyString, testKey)
			decrypted := Decrypt.StringToBytes(emptyEncrypted, testKey)
			Expect(decrypted).To(BeEmpty())
		})
	})

	Describe("StringToString", func() {
		It("should decrypt string to string correctly", func() {
			decrypted := Decrypt.StringToString(encryptedStr, testKey)
			Expect(decrypted).To(Equal(testString))
		})

		It("should handle empty input", func() {
			emptyEncrypted := Encrypt.StringToString(emptyString, testKey)
			decrypted := Decrypt.StringToString(emptyEncrypted, testKey)
			Expect(decrypted).To(BeEmpty())
		})
	})

	Describe("Base64ToBytes", func() {
		It("should decrypt base64 to bytes correctly", func() {
			decrypted := Decrypt.Base64ToBytes(base64Str, testKey)
			Expect(string(decrypted)).To(Equal(testString))
		})

		It("should handle empty input", func() {
			emptyEncrypted := Encrypt.StringToBase64(emptyString, testKey)
			decrypted := Decrypt.Base64ToBytes(emptyEncrypted, testKey)
			Expect(decrypted).To(BeEmpty())
		})
	})

	Describe("Base64ToString", func() {
		It("should decrypt base64 to string correctly", func() {
			decrypted := Decrypt.Base64ToString(base64Str, testKey)
			Expect(decrypted).To(Equal(testString))
		})

		It("should handle empty input", func() {
			emptyEncrypted := Encrypt.StringToBase64(emptyString, testKey)
			decrypted := Decrypt.Base64ToString(emptyEncrypted, testKey)
			Expect(decrypted).To(BeEmpty())
		})
	})
})
