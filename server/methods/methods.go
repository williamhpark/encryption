package methods

import (
	"crypto/aes"
	"encoding/hex"
)

// Caesar Cipher encryption/decryption
// Params:
// `input` - string to encrypt/decrypt
// `key` - number of letters to shift by
// `encrypt` - true to encrypt, false to decrypt
func Caesar(input string, key int, encrypt bool) string {
	// Cast the string to an array of runes
	runes := []rune(input)

	// Cast the key to a rune
	shift := rune(key)
	// Size of the character set (26 in this case)
	offset := rune(26)

	switch encrypt {
	// Encrypt
	case true:
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
	// Decrypt
	case false:
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
	default:
		panic("Unexpected value for `encrypt` argument")
	}

	return string(runes)
}

// Advanced Encryption Standard (AES) encryption/decryption
// Params:
// `input` - string to encrypt/decrypt
// `s_key` - cipher key
// `encrypt` - true to encrypt, false to decrypt
func AES(input string, s_key string, encrypt bool) string {
	// Cast the cipher key to an array of bytes
	key := []byte(s_key)

	switch encrypt {
	// Encrypt
	case true:
		// Create cipher
		c, err := aes.NewCipher(key)
		if err != nil {
			panic(err)
		}

		// Allocate space for ciphered data
		out := make([]byte, len(input))

		// Encrypt the string
		c.Encrypt(out, []byte(input))

		// Return the hexadecimal encoding for the string
		return hex.EncodeToString(out)
	// Decrypt
	case false:
		// Retrieve the bytes represented by the hexadecimal string
		ciphertext, _ := hex.DecodeString(input)

		// Create cipher
		c, err := aes.NewCipher(key)
		if err != nil {
			panic(err)
		}

		// Allocate space for ciphered data
		pt := make([]byte, len(ciphertext))

		// Decrypt the string
		c.Decrypt(pt, ciphertext)

		return string(pt[:])
	default:
		panic("Unexpected value for `encrypt` argument")
	}
}
