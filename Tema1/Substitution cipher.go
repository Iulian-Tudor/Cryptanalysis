package main

import (
	"fmt"
	"strings"
)

// Caesar cipher (Shifting) encryption
func CaesarEncrypt(plainText string, shift int) string {
	plainText = strings.ToUpper(plainText)
	var cipherText string
	for _, char := range plainText {
		if char >= 'A' && char <= 'Z' {
			cipherText += string(((char - 'A' + rune(shift)) % 26) + 'A')
		} else {
			cipherText += string(char) // for spaces and other characters
		}
	}
	return cipherText
}

// Caesar cipher (Shifting) decryption
func CaesarDecrypt(cipherText string, shift int) string {
	cipherText = strings.ToUpper(cipherText)
	var plainText string
	for _, char := range cipherText {
		if char >= 'A' && char <= 'Z' {
			plainText += string(((char - 'A' - rune(shift) + 26) % 26) + 'A')
		} else {
			plainText += string(char) // for spaces and other characters
		}
	}
	return plainText
}

func modInverse(a, m int) int {
	for x := 1; x < m; x++ {
		if (a*x)%m == 1 {
			return x
		}
	}
	return -1 // no modular inverse found
}

// Affine cipher encryption: E(x) = (a * x + b) % 26
func AffineEncrypt(plainText string, a, b int) string {
	plainText = strings.ToUpper(plainText)
	var cipherText string
	for _, char := range plainText {
		if char >= 'A' && char <= 'Z' {
			x := int(char - 'A') // Convert char to number (A=0, B=1, ..., Z=25)
			cipherChar := (a*x + b) % 26
			cipherText += string(rune(cipherChar + 'A'))
		} else {
			cipherText += string(char) // for spaces and other characters
		}
	}
	return cipherText
}

// Affine cipher decryption: D(x) = a_inverse * (x - b) % 26
func AffineDecrypt(cipherText string, a, b int) string {
	cipherText = strings.ToUpper(cipherText)
	a_inv := modInverse(a, 26)
	if a_inv == -1 {
		return "Invalid key for decryption: a and 26 are not coprime."
	}

	var plainText string
	for _, char := range cipherText {
		if char >= 'A' && char <= 'Z' {
			x := int(char - 'A') // Convert char to number (A=0, B=1, ..., Z=25)
			plainChar := (a_inv * (x - b + 26)) % 26
			plainText += string(rune(plainChar + 'A'))
		} else {
			plainText += string(char) // for spaces and other characters
		}
	}
	return plainText
}

func main() {
	// Caesar cipher (shifting)
	fmt.Println("Caesar Cipher (Shifting) Test:")
	testcase := "HELLO HOW ARE YOU"
	shift := 3
	caesarCipherText := CaesarEncrypt(testcase, shift)
	caesarPlainText := CaesarDecrypt(caesarCipherText, shift)
	fmt.Println("Plain Text: ", testcase)
	fmt.Println("Cipher Text (Shift 3): ", caesarCipherText)
	fmt.Println("Decrypted Text: ", caesarPlainText)

	// Affine cipher
	fmt.Println("\nAffine Cipher Test:")
	affineA := 5
	affineB := 8
	affineCipherText := AffineEncrypt(testcase, affineA, affineB)
	affinePlainText := AffineDecrypt(affineCipherText, affineA, affineB)
	fmt.Println("Plain Text: ", testcase)
	fmt.Println("Cipher Text: ", affineCipherText)
	fmt.Println("Decrypted Text: ", affinePlainText)
}
