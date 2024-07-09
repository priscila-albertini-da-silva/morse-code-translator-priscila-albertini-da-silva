package translator

import (
	"fmt"
	"strings"
)

var morseCodeMap = map[string]string{
	".-":    "A",
	"-...":  "B",
	"-.-.":  "C",
	"-..":   "D",
	".":     "E",
	"..-.":  "F",
	"--.":   "G",
	"....":  "H",
	"..":    "I",
	".---":  "J",
	"-.-":   "K",
	".-..":  "L",
	"--":    "M",
	"-.":    "N",
	"---":   "O",
	".--.":  "P",
	"--.-":  "Q",
	".-.":   "R",
	"...":   "S",
	"-":     "T",
	"..-":   "U",
	"...-":  "V",
	".--":   "W",
	"-..-":  "X",
	"-.--":  "Y",
	"--..":  "Z",
	"-----": "0",
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",
}

type Translator struct {
	morseToText map[string]string
}

func NewTranslator() *Translator {
	return &Translator{
		morseToText: morseCodeMap,
	}
}

func (t *Translator) TranslateMorseToText(input string) string {
	input = strings.TrimSpace(input)

	words := strings.Split(input, "   ")

	if t.isArrayEmpty(words) {
		fmt.Println("Empty input. Please enter morse code.")
		return ""
	}

	var result []string = t.decodeMorseWords(words)

	if t.isArrayEmpty(result) {
		fmt.Println("Invalid input. Please enter morse code.")
		return ""
	}

	return strings.Join(result, " ")
}

func (t *Translator) isArrayEmpty(array []string) bool {
	return len(array) == 0 || (len(array) == 1 && (array)[0] == "")
}

func (t *Translator) decodeMorseWords(words []string) []string {
	var result []string

	for _, word := range words {
		var decodedWord []string
		letters := strings.Split(word, " ")
		for _, letter := range letters {
			if decodedLetter, exists := t.morseToText[letter]; exists {
				decodedWord = append(decodedWord, decodedLetter)
			}
		}
		result = append(result, strings.Join(decodedWord, ""))
	}

	return result
}
