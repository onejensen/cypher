package cipher

import (
	"strings"
)

// EncryptROT13 implements the ROT13 encryption algorithm
// Input: input - string to be encrypted
// Output: encrypted string
// Working Principle: Shifts each letter in the string forward by 13 positions
func EncryptROT13(input string) string {
	return shiftRune(input, 13)
}

// DecryptROT13 implements the ROT13 decryption algorithm
// Input: input - string to be decrypted
// Output: decrypted string
// Working Principle: Shifts each letter in the string backward by 13 positions
func DecryptROT13(input string) string {
	return shiftRune(input, -13)
}

// EncryptReverse implements the Reverse encryption algorithm
// Input: input - string to be encrypted
// Output: encrypted string
// Working Principle: Reverses the order of characters in the string
func EncryptReverse(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// DecryptReverse implements the Reverse decryption algorithm
// Input: input - string to be decrypted
// Output: decrypted string
// Note: In Reverse encryption, encryption and decryption operations are the same
func DecryptReverse(input string) string {
	return EncryptReverse(input)
}

// shiftRune shifts each character in the string by the specified amount
// Input: input - string to be processed, shift - amount of shift
// Output: string processed with the shift
// Working Principle: Shifts letter characters based on the shift amount, non-letter characters remain unchanged
func shiftRune(input string, shift int) string {
	var result strings.Builder
	for _, r := range input {
		if r >= 'a' && r <= 'z' {
			result.WriteRune('a' + (r-'a'+rune(shift)+26)%26)
		} else if r >= 'A' && r <= 'Z' {
			result.WriteRune('A' + (r-'A'+rune(shift)+26)%26)
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// EncryptCaesar encrypts a message using the Caesar cipher
// Input: message - message to be encrypted, shift - amount of shift
// Output: encrypted message
// Working Principle: Shifts each letter character in the message based on the shift amount, non-letter characters remain unchanged
func EncryptCaesar(message string, shift int) string {
	result := ""
	for _, ch := range message {
		if ch >= 'a' && ch <= 'z' {
			shifted := (ch-'a'+rune(shift))%26 + 'a'
			result += string(shifted)
		} else if ch >= 'A' && ch <= 'Z' {
			shifted := (ch-'A'+rune(shift))%26 + 'A'
			result += string(shifted)
		} else {
			result += string(ch)
		}
	}
	return result
}

// DecryptCaesar decrypts a message using the Caesar cipher
// Input: message - message to be decrypted, shift - amount of shift
// Output: decrypted message
// Working Principle: Shifts each letter character in the message in the reverse direction based on the shift amount, non-letter characters remain unchanged
func DecryptCaesar(message string, shift int) string {
	result := ""
	for _, ch := range message {
		if ch >= 'a' && ch <= 'z' {
			shifted := (ch-'a'-rune(shift)+26)%26 + 'a'
			result += string(shifted)
		} else if ch >= 'A' && ch <= 'Z' {
			shifted := (ch-'A'-rune(shift)+26)%26 + 'A'
			result += string(shifted)
		} else {
			result += string(ch)
		}
	}
	return result
}
