package roman

import "strings"

// Version 1
// func ConvertToRoman(arabic int) string {
// 	if arabic == 1 {
// 		return "I"
// 	} else if arabic == 2 {
// 		return "II"
// 	} else {
// 		return "III"
// 	}
// }

// Version 2
// using StringBuilder
// func ConvertToRoman(arabic int) string {
// 	if arabic == 5 {
// 		return "V"
// 	}
// 	if arabic == 4 {
// 		return "IV"
// 	}

// 	var result strings.Builder

// 	for i := 0; i < arabic; i++ {
// 		result.WriteString("I")
// 	}

// 	return result.String()
// }

// Version 3, the if statements are running inside the loop
// GG la, lol
// func ConvertToRoman(arabic int) string {

// 	var result strings.Builder

// 	for i := arabic; i > 0; i-- {
// 		if i == 4 {
// 			result.WriteString("IV")
// 			break
// 		}
// 		result.WriteString("I")
// 	}

// 	return result.String()
// }

// Version4
// func ConvertToRoman(arabic int) string {
// 	var result strings.Builder
// 	for arabic > 0 {
// 		switch {
// 		case arabic > 9:
// 			result.WriteString("X")
// 			arabic -= 10
// 		case arabic > 8:
// 			result.WriteString("IX")
// 			arabic -= 9
// 		case arabic > 4:
// 			result.WriteString("V")
// 			arabic -= 5
// 		case arabic > 3:
// 			result.WriteString("IV")
// 			arabic -= 4
// 		default:
// 			result.WriteString("I")
// 			arabic--
// 		}
// 	}

// 	return result.String()
// }

type romanNumeral struct {
	Value  int
	Symbol string
}

type romanNumerals []romanNumeral

var allRomanNumerals = romanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (r romanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func (r romanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{byte(symbol), byte(w[i+1])})
			i++
		} else {
			symbols = append(symbols, []byte{byte(symbol)})
		}
	}
	return
}

func ConvertToArabic(roman string) (total int) {
	for _, symbols := range windowedRoman(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbols...)
	}
	return total
}
