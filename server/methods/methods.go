package methods

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// Caesar Cipher encryption
// Params:
// `message` - string to encrypt
// `key` - number of letters to shift by
func CaesarEncrypt(message string, key int) (string, error) {
	// Cast the string to an array of runes
	runes := []rune(message)
	// Cast the key to a rune
	shift := rune(key)

	// Size of the character set (26 in this case)
	offset := rune(26)

	for index, char := range runes {
		if char >= 'a'+shift && char <= 'z' ||
			char >= 'A'+shift && char <= 'Z' {
			char = char - shift
		} else if char >= 'a' && char < 'a'+shift ||
			char >= 'A' && char < 'A'+shift {
			char = char - shift + offset
		}
		runes[index] = char
	}

	return string(runes), nil
}

// Caesar Cipher decryption
// Params:
// `message` - string to decrypt
// `key` - number of letters to shift by
func CaesarDecrypt(message string, key int) (string, error) {
	// Cast the string to an array of runes
	runes := []rune(message)
	// Cast the key to a rune
	shift := rune(key)

	// Size of the character set (26 in this case)
	offset := rune(26)

	for index, char := range runes {
		if char >= 'a' && char <= 'z'-shift ||
			char >= 'A' && char <= 'Z'-shift {
			char = char + shift
		} else if char > 'z'-shift && char <= 'z' ||
			char > 'Z'-shift && char <= 'Z' {
			char = char + shift - offset
		}
		runes[index] = char
	}

	return string(runes), nil
}

// Advanced Encryption Standard (AES) encryption
// Params:
// `message` - string to encrypt
// `s_key` - cipher key
func AESEncrypt(message string, s_key string) (string, error) {
	// Cast the cipher key to an array of bytes
	key := []byte(s_key)

	// Create cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Allocate space for ciphered data
	out := make([]byte, len(message))

	// Encrypt the string
	c.Encrypt(out, []byte(message))

	// Return the hexadecimal encoding for the string
	return hex.EncodeToString(out), nil
}

// Advanced Encryption Standard (AES) decryption
// Params:
// `message` - string to decrypt
// `sKey` - cipher key
func AESDecrypt(message string, sKey string) (string, error) {
	// Cast the cipher key to an array of bytes
	key := []byte(sKey)

	// Retrieve the bytes represented by the hexadecimal string
	ciphertext, _ := hex.DecodeString(message)

	// Create cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Allocate space for ciphered data
	pt := make([]byte, len(ciphertext))

	// Decrypt the string
	c.Decrypt(pt, ciphertext)

	return string(pt[:]), nil
}

// Generate an RSA private key
func GenerateRSAPrivateKey() rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	return *privateKey
}

// Rivest-Shamir-Adleman (RSA) encryption
// Params:
// `message` - string to encrypt
// `key` - public key
func RSAEncrypt(message string, key rsa.PublicKey) (string, error) {
	label := []byte("OAEP Encrypted")
	// `crypto/rand.Reader`` is a good source of entropy for randomizing the encryption function
	rng := rand.Reader

	// Encrypts the given message with RSA-OAEP
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &key, []byte(message), label)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Rivest-Shamir-Adleman (RSA) encryption
// Params:
// `message` - string to encrypt/decrypt
// `key` - private key
func RSADecrypt(message string, key rsa.PrivateKey) (string, error) {
	ct, _ := base64.StdEncoding.DecodeString(message)
	label := []byte("OAEP Encrypted")
	rng := rand.Reader

	// Decrypt the ciphertext using RSA-OAEP
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, &key, ct, label)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
