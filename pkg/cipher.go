package cipher

import (
	"strings"
)

// EncryptROT13 实现ROT13加密算法
func EncryptROT13(input string) string {
	return shiftRune(input, 13)
}

// DecryptROT13 实现ROT13解密算法
func DecryptROT13(input string) string {
	return shiftRune(input, -13)
}

// EncryptReverse 实现Reverse加密算法
func EncryptReverse(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// DecryptReverse 实现Reverse解密算法
// 在Reverse加密中，加密和解密操作是相同的
func DecryptReverse(input string) string {
	return EncryptReverse(input)
}

// shiftRune 对字符串中的每个字符进行位移
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

// EncryptCaesar 加密消息使用凯撒密码
func EncryptCaesar(message string, shift int) string {
	result := ""
	for _, ch := range message {
		if ch >= 'a' && ch <= 'z' {
			// 对小写字母进行位移
			shifted := (ch-'a'+rune(shift))%26 + 'a'
			result += string(shifted)
		} else if ch >= 'A' && ch <= 'Z' {
			// 对大写字母进行位移
			shifted := (ch-'A'+rune(shift))%26 + 'A'
			result += string(shifted)
		} else {
			// 非字母字符保持不变
			result += string(ch)
		}
	}
	return result
}

// DecryptCaesar 解密消息使用凯撒密码
func DecryptCaesar(message string, shift int) string {
	result := ""
	for _, ch := range message {
		if ch >= 'a' && ch <= 'z' {
			// 对小写字母进行反向位移
			shifted := (ch-'a'-rune(shift)+26)%26 + 'a'
			result += string(shifted)
		} else if ch >= 'A' && ch <= 'Z' {
			// 对大写字母进行反向位移
			shifted := (ch-'A'-rune(shift)+26)%26 + 'A'
			result += string(shifted)
		} else {
			// 非字母字符保持不变
			result += string(ch)
		}
	}
	return result
}
