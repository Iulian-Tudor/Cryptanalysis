package main

import (
	"fmt"
	"strings"
)

// Constants
var englishFreq = map[rune]float64{
	'E': 0.1270, 'T': 0.0906, 'A': 0.0817, 'O': 0.0751, 'I': 0.0697, 'N': 0.0675,
	'S': 0.0633, 'H': 0.0609, 'R': 0.0599, 'D': 0.0425, 'L': 0.0403, 'C': 0.0278,
	'U': 0.0276, 'M': 0.0241, 'W': 0.0236, 'F': 0.0223, 'G': 0.0202, 'Y': 0.0197,
	'P': 0.0193, 'B': 0.0129, 'V': 0.0098, 'K': 0.0077, 'J': 0.0015, 'X': 0.0015,
	'Q': 0.0010, 'Z': 0.0007,
}

// Helper function to calculate absolute value of a float64
func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

// Remove all non-letters and convert to uppercase
func cleanText(text string) string {
	cleaned := ""
	for _, char := range strings.ToUpper(text) {
		if char >= 'A' && char <= 'Z' {
			cleaned += string(char)
		}
	}
	return cleaned
}

// Frequency analysis of text
func calculateLetterFrequencies(text string) map[rune]int {
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
	frequency := calculateLetterFrequencies(text)
	sum := 0.0

	for _, count := range frequency {
		sum += float64(count * (count - 1))
	}

	ic := sum / (n * (n - 1))
	return ic
}

// Find likely key length using Index of Coincidence method
func FindKeyLengthByIC(cipherText string, maxKeyLength int) int {
	cipherText = cleanText(cipherText)
	englishIC := 0.068 // Expected IC for English text
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

// Extract N-th subkey letters
func getNthSubkeysLetters(n, keyLength int, cipherText string) string {
	var subText string
	for i := n - 1; i < len(cipherText); i += keyLength {
		subText += string(cipherText[i])
	}
	return subText
}

// Shift the text by a certain number of positions in the alphabet
func shiftSubTextBy(text string, shift rune) string {
	shiftedText := ""
	for _, char := range text {
		shiftedText += string((char-'A'+shift)%26 + 'A')
	}
	return shiftedText
}

// Calculate IMC for a substring
func calculateIMC(subText string) float64 {
	subTextFreq := calculateLetterFrequencies(subText)
	var imc float64 = 0.0

	for letter, freq := range subTextFreq {
		imc += float64(freq) * englishFreq[letter] // Convert freq (int) to float64
	}

	return imc
}

// Extract key based on the IC method and frequency analysis
func extractKey(cipherText string, keyLength int) string {
	key := ""
	for i := 1; i <= keyLength; i++ {
		subText := getNthSubkeysLetters(i, keyLength, cipherText)
		maxIMC := 0.0
		var bestShift rune

		for shift := 'A'; shift <= 'Z'; shift++ {
			shiftedSubText := shiftSubTextBy(subText, shift)
			imc := calculateIMC(shiftedSubText)

			if imc > maxIMC {
				maxIMC = imc
				bestShift = shift
			}
		}

		key += string(bestShift)
	}

	return key
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
			shift := key[i%len(key)] - 'A'
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
			shift := key[i%len(key)] - 'A'
			plainChar := (cipherText[i] - 'A' - shift + 26) % 26 // Ensure non-negative result
			plainText += string(plainChar + 'A')
		}
	}
	return plainText
}

// Main test function
func main() {
	testcase := "HELLO HOW ARE YOU HOW ARE YOU DOING"
	key := "QWERTY"
	cipherText := EncryptVigenereCipher(testcase, key)
	plainText := DecryptVigenereCipher(cipherText, key)
	fmt.Println("Plain Text: ", plainText)
	fmt.Println("Cipher Text: ", cipherText)

	// Find key length using Index of Coincidence method
	keyLengthIC := FindKeyLengthByIC(cipherText, 10) // Assuming max possible key length is 10
	fmt.Printf("Likely key length from Index of Coincidence: %d\n", keyLengthIC)

	// Recover the key based on the detected key length
	recoveredKey := extractKey(cipherText, keyLengthIC)
	fmt.Printf("Recovered Key: %s\n", recoveredKey)
}
