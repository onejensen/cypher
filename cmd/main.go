// version 1.0
package main

import (
	"bufio"             // bufio package for reading user input
	cipher "cypher/pkg" // Custom cipher package containing encryption and decryption functions
	"cypher/util"       // util package for obtaining user operations and cipher types
	"fmt"               // For formatted output
	"os"                // For handling OS-level input and output
	"strings"           // For string processing
)

func main() {
	// Initialize a bufio.Reader for reading from standard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Cypher Tool!")

	for {
		// Get the user's choice of operation (encrypt or decrypt)
		toEncrypt, err := util.GetOperation(reader)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Get the user's choice of encryption type or to go back to the main menu
		encoding, backToMenu, err := util.GetCipherType(reader)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if backToMenu {
			continue
		}

		// Get the user's message
		fmt.Println("Enter the message:")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}
		// Remove whitespace from both ends of the message
		message = strings.TrimSpace(message)

		// Perform encryption or decryption based on the chosen type
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
		case "Caesar":
			if toEncrypt {
				result = cipher.EncryptCaesar(message, shift)
				operation = "Encryption"
			} else {
				result = cipher.DecryptCaesar(message, shift)
				operation = "Decryption"
			}
		}

		// Print the result of encryption or decryption
		fmt.Printf("%s result using %s:\n%s\n\n", operation, encoding, result)
		fmt.Println("Please save your message. \nPress Enter to continue...")

		// Wait for the user to press Enter
		_, _ = reader.ReadString('\n')

		// Ask user if they would like to perform another operation
		fmt.Println("Would you like to do another operation?")
		toChoose, err := util.GetOption(reader)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if !toChoose {
			fmt.Println("Thank you for using our app! (Created with love by Hao, Auri & Juanjo)")
			break
		}
	}
}
