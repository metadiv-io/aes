package aes

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Encrypt", func() {
	const testKey = "0123456789abcdef" // 16-byte key for AES-128

	Context("BytesToBytes", func() {
		It("should encrypt byte slice correctly", func() {
			input := []byte("hello world")
			encrypted := Encrypt.BytesToBytes(input, testKey)
			Expect(encrypted).NotTo(Equal(input))
			Expect(len(encrypted)).To(BeNumerically(">=", len(input)))
		})

		It("should handle empty input", func() {
			input := []byte("")
			encrypted := Encrypt.BytesToBytes(input, testKey)
			Expect(len(encrypted)).To(Equal(16)) // One AES block
		})
	})

	Context("BytesToString", func() {
		It("should encrypt bytes to string", func() {
			input := []byte("test message")
			encrypted := Encrypt.BytesToString(input, testKey)
			Expect(encrypted).NotTo(Equal("test message"))
		})
	})

	Context("BytesToBase64", func() {
		It("should encrypt bytes to base64 string", func() {
			input := []byte("hello world")
			encrypted := Encrypt.BytesToBase64(input, testKey)
			Expect(encrypted).NotTo(Equal("hello world"))
			// Base64 string should only contain valid characters
			Expect(encrypted).To(MatchRegexp("^[A-Za-z0-9+/]*={0,2}$"))
		})
	})

	Context("StringToBytes", func() {
		It("should encrypt string to bytes", func() {
			input := "hello world"
			encrypted := Encrypt.StringToBytes(input, testKey)
			Expect(encrypted).NotTo(Equal([]byte(input)))
			Expect(len(encrypted)).To(BeNumerically(">=", len(input)))
		})
	})

	Context("StringToString", func() {
		It("should encrypt string to string", func() {
			input := "test message"
			encrypted := Encrypt.StringToString(input, testKey)
			Expect(encrypted).NotTo(Equal(input))
		})
	})

	Context("StringToBase64", func() {
		It("should encrypt string to base64", func() {
			input := "hello world"
			encrypted := Encrypt.StringToBase64(input, testKey)
			Expect(encrypted).NotTo(Equal(input))
			// Base64 string should only contain valid characters
			Expect(encrypted).To(MatchRegexp("^[A-Za-z0-9+/]*={0,2}$"))
		})
	})
})
