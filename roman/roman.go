package roman

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

// slice declaration
type RomanNumerals []RomanNumeral

// slice
var romanNumerals = RomanNumerals{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

func (r RomanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

func ToRoman(arabic int) string {
	var result strings.Builder
	for _, numeral := range romanNumerals {
		for arabic >= numeral.Value {
			arabic -= numeral.Value
			result.WriteString(numeral.Symbol)
		}
	}
	return result.String()
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	isSubtractiveSymbol := currentSymbol == 'I'
	return index+1 < len(roman) && isSubtractiveSymbol
}

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractive(symbol) && romanNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{byte(symbol), byte(w[i+1])})
			i++
		} else {
			symbols = append(symbols, []byte{byte(symbol)})
		}
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I'
}

func ToArabic(roman string) (total int) {
	for _, symbols := range windowedRoman(roman).Symbols() {
		total += romanNumerals.ValueOf(symbols...)
	}
	return
}
