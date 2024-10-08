package main

import (
	"fmt"
	"sort"
	"strings"
)

// Frecvențele literelor în limba engleză
var englishFrequencies = []rune{'E', 'T', 'A', 'O', 'I', 'N', 'S', 'H', 'R', 'D', 'L', 'C', 'U', 'M', 'W', 'F', 'G', 'Y', 'P', 'B', 'V', 'K', 'J', 'X', 'Q', 'Z'}

// Substitution cipher key provided by the user
var substitutionKey = map[rune]rune{
	'A': 'O', 'B': 'P', 'C': 'Q', 'D': 'R', 'E': 'S', 'F': 'T', 'G': 'U', 'H': 'V',
	'I': 'W', 'J': 'X', 'K': 'Y', 'L': 'Z', 'M': 'A', 'N': 'B', 'O': 'C', 'P': 'D',
	'Q': 'E', 'R': 'F', 'S': 'G', 'T': 'H', 'U': 'I', 'V': 'J', 'W': 'K', 'X': 'L',
	'Y': 'M', 'Z': 'N',
}

// Inversarea cheii pentru decriptare
var reverseSubstitutionKey = map[rune]rune{}

func init() {
	// Generăm cheia inversă pentru decriptare
	for k, v := range substitutionKey {
		reverseSubstitutionKey[v] = k
	}
}

// Functie pentru criptarea textului folosind o cheie de substituție
func encrypt(plaintext string, key map[rune]rune) string {
	encrypted := []rune{}
	for _, char := range strings.ToUpper(plaintext) {
		if newChar, ok := key[char]; ok {
			encrypted = append(encrypted, newChar)
		} else {
			encrypted = append(encrypted, char) // Dacă nu este o literă, lăsăm caracterul neschimbat
		}
	}
	return string(encrypted)
}

// Functie pentru calcularea frecvențelor din text
func calculateFrequency(text string) map[rune]float64 {
	frequency := make(map[rune]float64)
	total := 0

	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			frequency[char]++
			total++
		}
	}

	// Calculăm procentajele
	for char := range frequency {
		frequency[char] = (frequency[char] / float64(total)) * 100
	}

	return frequency
}

// Functie pentru compararea frecventelor și sortarea lor
func compareFrequencies(ciphertext string) []rune {
	ciphertext = strings.ToUpper(ciphertext)
	freq := calculateFrequency(ciphertext)

	// Sortare după frecvență
	type pair struct {
		char  rune
		value float64
	}
	var freqPairs []pair
	for k, v := range freq {
		freqPairs = append(freqPairs, pair{k, v})
	}
	sort.Slice(freqPairs, func(i, j int) bool {
		return freqPairs[i].value > freqPairs[j].value
	})

	fmt.Println("Frecvențele în criptotext:")
	for _, p := range freqPairs {
		fmt.Printf("%c: %.2f%%\n", p.char, p.value)
	}

	// Returnăm literele sortate în funcție de frecvență
	var sortedCipherLetters []rune
	for _, p := range freqPairs {
		sortedCipherLetters = append(sortedCipherLetters, p.char)
	}
	return sortedCipherLetters
}

// Functie pentru generarea unei chei pe baza frecvențelor
func generateSubstitutionKey(cipherLetters []rune) map[rune]rune {
	// Generăm o mapare între literele din criptotext și literele din limba engleză în ordinea frecvenței
	substitutionKey := make(map[rune]rune)
	for i := 0; i < len(cipherLetters) && i < len(englishFrequencies); i++ {
		substitutionKey[cipherLetters[i]] = englishFrequencies[i]
	}

	return substitutionKey
}

// Functie pentru aplicarea substituției și decriptarea textului
func decryptWithKey(ciphertext string, key map[rune]rune) string {
	decrypted := []rune{}
	for _, char := range strings.ToUpper(ciphertext) {
		if newChar, ok := key[char]; ok {
			decrypted = append(decrypted, newChar)
		} else {
			decrypted = append(decrypted, char) // În cazul în care nu găsim o substituție
		}
	}
	return string(decrypted)
}

func main() {
	// Introducerea textului original
	plaintext := "HELLO WORLD THIS IS A SECRET MESSAGE."

	// Criptăm textul
	encryptedText := encrypt(plaintext, substitutionKey)
	fmt.Println("Text criptat:", encryptedText)

	// Comparăm frecvențele și obținem literele sortate în funcție de frecvență
	cipherLetters := compareFrequencies(encryptedText)

	// Generăm o cheie pe baza frecvențelor
	generatedKey := generateSubstitutionKey(cipherLetters)

	// Afișăm cheia generată
	fmt.Println("Cheia de substituție generată din frecvențe:")
	for cipherChar, plainChar := range generatedKey {
		fmt.Printf("%c -> %c\n", cipherChar, plainChar)
	}

	// Decriptăm criptotextul folosind cheia generată din frecvențe
	decryptedText := decryptWithKey(encryptedText, generatedKey)
	fmt.Println("Text decriptat (bazat pe analiză de frecvență):", decryptedText)

	// Decriptăm criptotextul folosind cheia reală (cea inversată)
	actualDecryptedText := decryptWithKey(encryptedText, reverseSubstitutionKey)
	fmt.Println("Text decriptat corect (cu cheia reală):", actualDecryptedText)
}
