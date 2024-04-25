# AES Encryption

## Installation

```base
go get -u github.com/metadiv-io/aes
```

### String Encryption

Encrypt a string to bytes:

```go
b := aes.EncryptStringToBytes("Hello, World!", "encrypt key")
```

Encrypt a string to string

```go
s := aes.EncryptStringToString("Hello, World!", "encrypt key")
```

Encrypt a string to base64 string

```go
s := aes.EncryptStringToBase64("Hello, World!", "encrypt key")
```

### String Decryption

Decrypt a string from bytes:

```go
s := aes.DecryptBytesToString(b, "encrypt key")
```

Decrypt a string from string:

```go
s := aes.DecryptStringToString(s, "encrypt key")
```

Decrypt a string from base64 string:

```go
s := aes.DecryptBase64ToString(s, "encrypt key")
```

### Encrypt Bytes

Encrypt bytes to bytes:

```go
b := aes.EncryptBytesToBytes([]byte("Hello, World!"), "encrypt key")
```

Encrypt bytes to string:

```go
s := aes.EncryptBytesToString([]byte("Hello, World!"), "encrypt key")
```

Encrypt bytes to base64 string:

```go
s := aes.EncryptBytesToBase64([]byte("Hello, World!"), "encrypt key")
```

### Decrypt Bytes

Decrypt bytes from bytes:

```go
b := aes.DecryptBytesToBytes(b, "encrypt key")
```

Decrypt bytes from string:

```go
b := aes.DecryptStringToBytes(s, "encrypt key")
```

Decrypt bytes from base64 string:

```go
b := aes.DecryptBase64ToBytes(s, "encrypt key")
```
