package translator

import (
	"testing"
)

func TestTranslateMorseToText(t *testing.T) {
	tr := NewTranslator()
	tests := []struct {
		input    string
		expected string
	}{
		{".... . -.--   .--- ..- -.. .", "HEY JUDE"},
		{".-.. . -   .. -   -... .", "LET IT BE"},
		{"... --- ...", "SOS"},
		{"- .... .   --.- ..- .. -.-. -.-   -... .-. --- .-- -.   ..-. --- -..-", "THE QUICK BROWN FOX"},
		{".---- ..--- ...-- ....- ..... -.... --... ---.. ----. -----", "1234567890"},
		{"   .... . .-.. .-.. ---   .-- --- .-. .-.. -..   ", "HELLO WORLD"},
		{".... . .-.. .-.. ---", "HELLO"},
		{"-- --- .-. ... .", "MORSE"},
		{"", ""},
		{"   ", ""},
		{"ERROR_TEST", ""},
	}

	for _, test := range tests {
		output := tr.TranslateMorseToText(test.input)
		if output != test.expected {
			t.Errorf("TranslateMorseToText(%q) = %q; want %q", test.input, output, test.expected)
		}
	}
}
