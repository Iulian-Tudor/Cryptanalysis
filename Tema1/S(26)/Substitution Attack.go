package main

import (
	"fmt"
	"sort"
	"strings"
)

var englishFrequencies = map[rune]float64{
	'E': 12.702, 'T': 9.056, 'A': 8.167, 'O': 7.507, 'I': 6.966, 'N': 6.749,
	'S': 6.327, 'H': 6.094, 'R': 5.987, 'D': 4.253, 'L': 4.025, 'C': 2.782,
	'U': 2.758, 'M': 2.406, 'W': 2.360, 'F': 2.228, 'G': 2.015, 'Y': 1.974,
	'P': 1.929, 'B': 1.492, 'V': 0.978, 'K': 0.772, 'J': 0.153, 'X': 0.150,
	'Q': 0.095, 'Z': 0.074,
}

func calculateFrequency(text string) map[rune]float64 {
	frequency := make(map[rune]float64)
	total := 0

	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			frequency[char]++
			total++
		}
	}

	// Calculate percentages
	for char := range frequency {
		frequency[char] = (frequency[char] / float64(total)) * 100
	}

	return frequency
}

func compareFrequencies(ciphertext string) {
	ciphertext = strings.ToUpper(ciphertext)
	freq := calculateFrequency(ciphertext)

	// Sort by frequency
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

	fmt.Println("Frequencies in ciphertext:")
	for _, p := range freqPairs {
		fmt.Printf("%c: %.2f%%\n", p.char, p.value)
	}
}

func main() {
	// Exemplu de criptotext
	ciphertext := "HVS OCKF HGSH!"

	// Comparare frecvențe
	compareFrequencies(ciphertext)
}
