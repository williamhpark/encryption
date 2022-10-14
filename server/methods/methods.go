package methods

// Caesar Cipher encryption/decryption
// Params:
// `input` - string to encrypt/decrypt
// `key` - number of letters to shift by
// `direction` - true to encrypt, false to decrypt
func Caesar(input string, key int, encrypt bool) string {
	// Cast the string to an array of runes
	runes := []rune(input)

	// Cast the key to a rune
	shift := rune(key)
	// Size of the character set (26 in this case)
	offset := rune(26)

	switch encrypt {
	// Decoding
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
	// Encoding
	default:
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
	}

	return string(runes)
}
