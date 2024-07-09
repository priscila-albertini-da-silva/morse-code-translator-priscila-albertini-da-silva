package translator

type TranslatorInterface interface {
	TranslateMorseToText(input string) string
}
