# aes

## Installation

```bash
go get -u github.com/metadiv-io/aes
```

## Highlights

## bytes <> string

* aes.EncryptToBytes(text, key string) []byte

* aes.DecryptFromBytes(encrypted []byte, key string) string

## bytes <> bytes

* aes.EncryptBytesToBytes(src []byte, key string) []byte

* aes.DecryptBytesFromBytes(encrypted []byte, key string) []byte

## base64 <> string

* aes.EncryptToBase64(text, key string) string

* DecryptFromBase64(encrypted string, key string) string
