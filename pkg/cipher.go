package cipher

import (
	"strings"
)

// EncryptROT13 applies the ROT13 encryption algorithm to the input string.
// It shifts each letter in the string forward by 13 positions.
func EncryptROT13(input string) string {
	return shiftRune(input, 13)
}

// DecryptROT13 applies the ROT13 decryption algorithm to the input string.
// It shifts each letter in the string backward by 13 positions.
func DecryptROT13(input string) string {
	return shiftRune(input, -13)
}

// EncryptReverse applies the Reverse encryption algorithm to the input string.
// It reverses the order of characters in the string.
func EncryptReverse(input string) string {
	var reversed string
	for i := len(input) - 1; i >= 0; i-- {
		reversed += string(input[i])
	}
	return reversed
}

// DecryptReverse applies the Reverse decryption algorithm to the input string.
// Note: In Reverse encryption, encryption and decryption operations are the same.
func DecryptReverse(input string) string {
	return EncryptReverse(input)
}

// shiftRune shifts each character in the string by the specified amount.
// It shifts letter characters based on the shift amount, non-letter characters remain unchanged.
func shiftRune(input string, shift int) string {
	var result strings.Builder
	for _, r := range input {
		if isLetter(r) {
			base := 'a'
			if isUpperCase(r) {
				base = 'A'
			}
			result.WriteRune(base + (r-base+rune(shift)+26)%26)
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// EncryptCaesar encrypts a message using the Caesar cipher.
// It shifts each letter character in the message based on the shift amount, non-letter characters remain unchanged.
func EncryptCaesar(message string, shift int) string {
	return shiftRune(message, shift)
}

// DecryptCaesar decrypts a message using the Caesar cipher.
/* It shifts each letter character in the message in the reverse direction based on the shift amount, non-letter
characters remain unchanged.*/
func DecryptCaesar(message string, shift int) string {
	return shiftRune(message, -shift)
}

// isLetter checks if a given rune is a letter.
func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// isUpperCase checks if a given rune is an uppercase letter.
func isUpperCase(r rune) bool {
	return r >= 'A' && r <= 'Z'
}
