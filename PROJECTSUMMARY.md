version 1.0
# Project Summary

## Project Directory Structure

### `main.go`
- **Description**: This is the entry file of the project.
  - Responsible for creating the user interface, processing user input, and calling the appropriate encryption or decryption methods.
  - Imports the custom cipher and util packages.

### `cipher` Package (`/cypher/pkg`)
- **Description**: Contains various encryption and decryption algorithms.
- **Functions**:
  - `EncryptROT13` and `DecryptROT13`: Implements the ROT13 algorithm.
  - `EncryptReverse` and `DecryptReverse`: Implements encryption and decryption of reversed strings.
  - `EncryptCaesar` and `DecryptCaesar`: Implements encryption and decryption using the Caesar cipher.

### `util` Package (`/cypher/util`)
- **Description**: Contains auxiliary functions to obtain the user's chosen operation (encryption/decryption) and the selected type of encryption (such as ROT13, Reverse, Caesar).

## 2. Code Logic and Execution Order

### Program Start
1. The user runs `main.go`, and the program starts executing.
2. Displays a welcome message and waits for user input.

### Selecting Operation Type (`util.GetOperation`)
1. The user is prompted to choose between encryption and decryption.
2. Inputting “1” represents encryption, while “2” represents decryption.
3. The choice is used to determine which encryption or decryption function to call later.

### Selecting Cipher Type (`util.GetCipherType`)
1. The user is prompted to choose an encryption algorithm: ROT13, Reverse, or Caesar.
2. Inputting “1” selects ROT13, “2” selects Reverse, and “3” selects Caesar Cipher.
3. The choice determines which encryption or decryption function to use subsequently.

### Inputting Message
1. The user inputs the message they want to encrypt or decrypt.
2. The input message is used for encryption or decryption processing.

### Executing Encryption or Decryption
1. Calls the respective encryption or decryption function based on the user's choice.
2. For Caesar Cipher, an additional shift amount is required.

### Detailed Operations
- **ROT13**: Calls `cipher.EncryptROT13` or `cipher.DecryptROT13`.
- **Reverse**: Both encryption and decryption call `cipher.EncryptReverse`.
- **Caesar Cipher**: Calls `cipher.EncryptCaesar` or `cipher.DecryptCaesar`, passing in the shift amount.

### Displaying Results
1. Displays the message after encryption or decryption.
2. The result depends on the user's operation type and the chosen algorithm.

### Program Termination
1. Prompts the user to press the Enter key to exit.
2. The program ends after the user responds.

## Encryption Algorithms

### 1. ROT13
- **Principle**: ROT13 is a simple substitution cipher that replaces each letter in the alphabet with the letter 13 positions after it in the alphabet. Since there are 26 letters in the English alphabet, this substitution is symmetric, meaning the same operation is used for both encryption and decryption.
- **Implementation**: The program checks each character. If the character is a letter, it shifts its ASCII value forward or backward by 13 places. For example, 'A' (ASCII 65) is replaced with 'N' (ASCII 78), and vice versa.

### 2. Reverse
- **Principle**: The Reverse algorithm encrypts by reversing the order of characters in a string. Since the reversing operation is its own inverse, the same method used for encryption can be used for decryption.
- **Implementation**: Implementing this algorithm involves iterating through the characters of the string, swapping the first character with the last, then the second with the second-to-last, and so on until all characters have been processed.

### 3. Caesar Cipher
- **Principle**: The Caesar cipher is a substitution encryption technique that encrypts by shifting the letters in the alphabet a fixed number of positions to the right. The decryption process involves shifting these letters the same number of positions to the left. The number of positions shifted is the key to the encryption.
- **Implementation**: When implementing the Caesar cipher, the program shifts each letter according to the shift amount (shift) provided by the user. For example, if the shift is 3, then 'A' is replaced with 'D', 'B' with 'E', and so on. For decryption, the shift amount is simply applied in the opposite direction. The program also needs to handle the wrapping of the alphabet to ensure that the shifting is cyclical (i.e., 'Z' is followed by 'A').
