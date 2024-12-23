package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {}

var symbols = generateSymbols()
var franMaxDigits = 8
var franPrintAll = false
var franPrintBeginAndEnd = true

func generateSymbols() []string {
	symbols := []string{}

	for i := 1; i <= 9; i++ {
		symbols = append(symbols, fmt.Sprintf("%d", i))
	}

	for i := 'a'; i <= 'z'; i++ {
		symbols = append(symbols, string(i))
	}

	for i := 'A'; i <= 'Z'; i++ {
		symbols = append(symbols, string(i))
	}

	specialChars := []string{"@", "#", "$", "%"}
	symbols = append(symbols, specialChars...)

	emojis := []string{"😊", "😍", "😎", "😭", "😡", "😱", "👍", "👏", "💪", "🎉", "🔥", "⭐"}
	symbols = append(symbols, emojis...)

	return symbols
}

func next(charBaseFran string) string {
	for i, char := range symbols {
		if char == charBaseFran {
			if i < len(symbols)-1 {
				return symbols[i+1]
			}
			return symbols[0]
		}
	}
	return ""
}

// Base Fran is an infinite-base number system. While there are an infinite amount of digits,
// each base-fran number is represented by a string of 5 digits. It's sort of like binary,
// but instead of wrapping to 100 after 11, it goes to 111. Then 1111, 11111, and finally 2.
// Then 21, 211, 2111, 21111, 22. 221, 2211, 22111, 222, 2221, 22211, 2222, 22221, 22222, 3.
func toBaseFran(nBase10 int) string {

	// This algorithm works by building the base-fran number
	// 1 by 1, until it gets to the value of nBase10.
	//
	// curr is the current base-fran number being built.
	curr := ""

	// How many numbers does it take from [x] to [xxxxx].
	// Resets when curr reaches 1 digit again.
	//
	// Used when franPrintBeginAndEnd is true.
	currFirstDigitCount := 0

	// Goes from nBase10 to 0 in steps of -1.
	// This 'algorithm' builds the Base Fran number 1 by 1.
	for i := 0; i < nBase10; i++ {
		currLen := len(curr)

		if franPrintBeginAndEnd {
			// If the number is single digit, we print the beginning of
			// the debug message.
			if currLen == 1 {
				currFirstDigitCount = 0
				fullFranOfSameDigit := strings.Repeat(curr, franMaxDigits)
				fmt.Printf("[%s] -> [%s] = %d ", curr, fullFranOfSameDigit, i)
			}

			// If it's 5 digits and they're all the same, we print the end.
			if currLen == franMaxDigits {
				onlyOneDigit := true
				for j := 0; j < franMaxDigits-1; j++ {
					if curr[j] != curr[j+1] {
						onlyOneDigit = false
						break
					}
				}
				if onlyOneDigit {
					fmt.Printf("-> %d. Total numbers for digit %s = %d.\n", i, string(curr[0]), currFirstDigitCount+1)
				}
			}
		}

		if franPrintAll {
			fmt.Printf("[%s] = %d\n", curr, i)
		}

		// If the number is less than 5 digits, for example '3332',
		// then we just add an '1' to the end.
		if currLen < franMaxDigits {
			curr += symbols[0]
			currFirstDigitCount++
			continue
		}

		// If it's a full number, we get the last digit.
		// Could be 1, 2, 3...
		lastChar := curr[franMaxDigits-1]

		// We delete all occurrences of that digit.
		// 33321 -> 3332
		// 33111 -> 33
		// 21111 -> 2
		// 11111 -> ""
		curr = strings.ReplaceAll(curr, string(lastChar), "")

		// And we append to the right the next digit in
		// the sequence, the digit with a value of 1 higher
		// than the last digit.
		curr += next(string(lastChar))
		currFirstDigitCount++
	}

	return curr
}
