package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/priscila-albertini-silva/morse-code-translator-priscila-albertini-da-silva/translator"
)

func main() {
	var ti translator.TranslatorInterface = translator.NewTranslator()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter Morse Code: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		output := ti.TranslateMorseToText(input)
		fmt.Println(output)
	}
}
