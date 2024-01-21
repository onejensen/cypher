package util

import (
	"bufio"   // bufio package for reading user input
	"fmt"     // For formatted output
	"strings" // For string processing
)

// GetOperation repeatedly gets user input until the user selects a valid operation (encrypt or decrypt)
// Input: reader - bufio.Reader object for reading from standard input
// Output: (bool, error) - The first return value is true for encrypt, false for decrypt. The second return value is for any possible errors.
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
			return true, nil // User chose encryption
		case "2":
			return false, nil // User chose decryption
		default:
			fmt.Printf("Invalid operation, please enter 1 or 2.\n")
		}
	}
}

// GetCipherType repeatedly gets user input until the user selects a valid cipher type or chooses to go back
// Input: reader - bufio.Reader object for reading from standard input
// Output: (string, bool, error) - The first return value is the user's selected cipher type ("ROT13", "Reverse", "Caesar"). The second return value is true if the user wants to go back to the main menu. The third return value is for any possible errors.
func GetCipherType(reader *bufio.Reader) (string, bool, error) {
	for {
		fmt.Printf("Select cipher (1/2/3/4):\n1. ROT13 (ONLY LATIN ALPHABET).\n2. Reverse.\n3. Caesar Cipher (ONLY LATIN ALPHABET).\n4. Back to Main Menu.\n")
		cipher, err := reader.ReadString('\n')
		if err != nil {
			return "", false, err
		}
		cipher = strings.TrimSpace(cipher)

		switch cipher {
		case "1":
			return "ROT13", false, nil // User chose ROT13 encryption
		case "2":
			return "Reverse", false, nil // User chose Reverse encryption
		case "3":
			return "Caesar", false, nil // User chose Caesar Cipher
		case "4":
			return "", true, nil // User chose to go back to the main menu
		default:
			fmt.Printf("Invalid cipher type, please enter 1, 2, 3, or 4.\n")
		}
	}
}

func GetOption(reader *bufio.Reader) (bool, error) {
	for {
		fmt.Printf("Select option (1/2):\n1. Yes.\n2. No.\n")
		opt, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}
		opt = strings.TrimSpace(opt)

		switch opt {
		case "1":
			return true, nil // User chose Yes
		case "2":
			return false, nil // User chose No
		default:
			fmt.Printf("Invalid option, please enter 1 or 2.\n")
		}
	}
}
