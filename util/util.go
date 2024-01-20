package util

import (
	"bufio"
	"fmt"
	"strings"
)

// GetOperation 修改后的函数，循环直到获得有效输入
func GetOperation(reader *bufio.Reader) (bool, error) {
	for {
		fmt.Printf("Select operation (1/2):\n1. Encrypt.\n2. Decrypt.\n")
		op, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}
		op = strings.TrimSpace(op)

		switch op {
		case "1":
			return true, nil // 表示加密
		case "2":
			return false, nil // 表示解密
		default:
			fmt.Printf("Invalid operation, please enter 1 or 2.\n")
		}
	}
}

// GetCipherType 修改后的函数，循环直到获得有效输入
func GetCipherType(reader *bufio.Reader) (string, error) {
	for {
		fmt.Printf("Select cipher (1/3):\n1. ROT13.\n2. Reverse.\n3. Caesar Cipher\n")
		cipher, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		cipher = strings.TrimSpace(cipher)

		switch cipher {
		case "1":
			return "ROT13", nil
		case "2":
			return "Reverse", nil
		case "3":
			return "Caesar", nil
		default:
			fmt.Printf("Invalid cipher type, please enter 1, 2, or 3.\n")
		}
	}
}
