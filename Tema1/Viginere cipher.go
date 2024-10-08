package main

import (
	"fmt"
	"strings"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Kasiski function to find key length using Kasiski method
func Kasiski(cipherText string) int {
	cipherText = strings.ReplaceAll(strings.ToUpper(cipherText), " ", "") // remove spaces
	sequenceLength := 3
	sequencePositions := make(map[string][]int)

	// Find all repeating sequences of length sequenceLength
	for i := 0; i < len(cipherText)-sequenceLength; i++ {
		sequence := cipherText[i : i+sequenceLength]
		sequencePositions[sequence] = append(sequencePositions[sequence], i)
	}

	var distances []int
	// For each sequence, find the distances between occurrences
	for _, positions := range sequencePositions {
		if len(positions) > 1 {
			for j := 1; j < len(positions); j++ {
				distances = append(distances, positions[j]-positions[j-1])
			}
		}
	}

	if len(distances) == 0 {
		fmt.Println("No repeating sequences found. Try longer text.")
		return -1
	}

	// Compute the GCD of the distances
	keyLength := distances[0]
	for _, d := range distances[1:] {
		keyLength = gcd(keyLength, d)
	}

	return keyLength
}

func letterFrequency(text string) map[rune]int {
	frequency := make(map[rune]int)
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			frequency[char]++
		}
	}
	return frequency
}

func IndexOfCoincidence(text string) float64 {
	text = strings.ReplaceAll(strings.ToUpper(text), " ", "") // remove spaces
	n := float64(len(text))
	frequency := letterFrequency(text)
	sum := 0.0

	for _, count := range frequency {
		sum += float64(count * (count - 1))
	}

	ic := sum / (n * (n - 1))
	return ic
}

func FindKeyLengthByIC(cipherText string, maxKeyLength int) int {
	cipherText = strings.ReplaceAll(strings.ToUpper(cipherText), " ", "") // remove spaces
	englishIC := 0.068                                                    // Expected IC for English text
	var bestKeyLength int
	var bestICDifference = 1.0

	for keyLength := 1; keyLength <= maxKeyLength; keyLength++ {
		var icSum float64

		// Split ciphertext into different groups based on key length
		for i := 0; i < keyLength; i++ {
			var subText string
			for j := i; j < len(cipherText); j += keyLength {
				subText += string(cipherText[j])
			}
			icSum += IndexOfCoincidence(subText)
		}

		averageIC := icSum / float64(keyLength)
		icDifference := abs(averageIC - englishIC)

		if icDifference < bestICDifference {
			bestICDifference = icDifference
			bestKeyLength = keyLength
		}
	}

	return bestKeyLength
}

// Helper function to calculate absolute value of a float64
func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

// EncryptVigenereCipher function to encrypt the plaintext using the Vigenère cipher
func EncryptVigenereCipher(plainText string, key string) string {
	plainText = strings.ToUpper(plainText)
	key = strings.ToUpper(key)
	var cipherText string
	for i := 0; i < len(plainText); i++ {
		if plainText[i] == 32 { // if it's a space, just add it
			cipherText += " "
		} else {
			// Shift the plaintext character by the key character
			shift := key[i%len(key)] - 'A' // Find the shift value from the key
			cipherChar := (plainText[i] - 'A' + shift) % 26
			cipherText += string(cipherChar + 'A')
		}
	}
	return cipherText
}

// DecryptVigenereCipher function to decrypt the cipher text using the Vigenère cipher
func DecryptVigenereCipher(cipherText string, key string) string {
	cipherText = strings.ToUpper(cipherText)
	key = strings.ToUpper(key)
	var plainText string
	for i := 0; i < len(cipherText); i++ {
		if cipherText[i] == 32 { // if it's a space, just add it
			plainText += " "
		} else {
			// Reverse the shift using the key
			shift := key[i%len(key)] - 'A'
			plainChar := (cipherText[i] - 'A' - shift + 26) % 26 // Ensure non-negative result
			plainText += string(plainChar + 'A')
		}
	}
	return plainText
}

// test cases
func main() {
	testcase := "HELLO HOW ARE YOU HOW ARE YOU DOING"
	key := "QWERTY"
	cipherText := EncryptVigenereCipher(testcase, key)
	plainText := DecryptVigenereCipher(cipherText, key)
	print("Plain Text: ", plainText, "\n")
	print("Cipher Text: ", cipherText, "\n")

	// Find key length using Kasiski's method
	keyLengthKasiski := Kasiski(cipherText)
	fmt.Printf("Likely key length from Kasiski's method: %d\n", keyLengthKasiski)

	// Find key length using Index of Coincidence method
	keyLengthIC := FindKeyLengthByIC(cipherText, 10) // Assuming max possible key length is 10
	fmt.Printf("Likely key length from Index of Coincidence: %d\n", keyLengthIC)
}
