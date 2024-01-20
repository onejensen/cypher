# Cypher

To run this program:

first type 
```sh 
cd sprint/cypher/cmd/
```
then: 
```sh
go run main.go
```

## Features

encrypt and decrypt text to:
- ROT13
- Reverse
- Caesar Cipher

## How to use the program:

After starting the program you can choose to either encrypt or decrypt text by selecting
"1" for Encrypt
"2" for Decrypt

After the selection you can select your encryption/decryption type:
"1" for ROT13
"2" for Reverse
"3" for Ceasar Cipher

After this selection:
Type the string of letters you wish to encrypt/decrypt.
For example encryption in ROT13:
```sh
Hello, World!
```
In ROT13 turns in to:
```sh
Uryyb, Jbeyq!
```
## The encryption methods explained:

## ROT13

ROT13 works by shifting the input letter 13 letters in the alphabet.

N O P Q R S T U V W X Y Z A B C D E F G H I J K L M 
^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
A B C D E F G H I J K L M N O P Q R S T U V W X Y Z

for example the word: "PUNY" becomes "CHAL" 
and the other way around "CHAL" becomes "PUNY"

## Reverse 

The Reverse Cypher works simply by reversing the word.

```sh
Hello, World!
```

Becomes: 
```sh
!dlroW ,olleH
```

## Caesar Cypher

The Caesar Cypher works similarly to the ROT13, but instead of shifting the letter by 13, it shifts the letter by x letters to the right (or left by inputting negative numbers).

With +3 the cypher works like this:

D E F G H I J K L M N O P Q R S T U V W X Y Z A B C
^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 

Example with cypher key "+3"
```sh
Hello, World!
```
Becomes:
```sh
Khoor, Zruog!
```


