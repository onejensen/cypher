package main

import (
	"bufio"
	cipher "cypher/pkg"
	"cypher/util"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Cypher Tool!")
	toEncrypt, err := util.GetOperation(reader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 获取用户选择的加密类型
	encoding, err := util.GetCipherType(reader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 获取用户的消息
	fmt.Println("Enter the message:")
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading message:", err)
		return
	}
	message = strings.TrimSpace(message)

	// 执行加密或解密
	var result, operation string
	var shift int
	if encoding == "Caesar" {
		fmt.Print("Enter the shift amount: ")
		fmt.Scanf("%d\n", &shift)
	}
	switch encoding {
	case "ROT13":
		if toEncrypt {
			result = cipher.EncryptROT13(message)
			operation = "Encryption"
		} else {
			result = cipher.DecryptROT13(message)
			operation = "Decryption"
		}
	case "Reverse":
		if toEncrypt {
			result = cipher.EncryptReverse(message)
			operation = "Encryption"
		} else {
			result = cipher.DecryptReverse(message)
			operation = "Decryption"
		}
		// 您可以在这里添加更多的加密方法
	case "Caesar":
		if toEncrypt {
			result = cipher.EncryptCaesar(message, shift)
			operation = "Encryption"
		} else {
			result = cipher.DecryptCaesar(message, shift)
			operation = "Decryption"
		}
	}

	fmt.Printf("%s result using %s: %s\n", operation, encoding, result)

	// 等待用户输入，以避免程序立即退出
	fmt.Println("Press 'Enter' to exit...")
	_, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
	}
}
