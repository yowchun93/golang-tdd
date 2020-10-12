package roman

import "strings"

// ConvertToRoman converts ArabicToRoman
// func ConvertToRoman(num int) string {
// 	if num == 1 {
// 		return "I"
// 	} else if num == 2 {
// 		return "II"
// 	}
// 	return "III"
// }

// ConvertToRoman converts ArabicToRoman
// uses loop
// func ConvertToRoman(num int) string {
// 	if num == 4 {
// 		return "IV"
// 	} else if num == 5 {
// 		return "V"
// 	}

// 	rom := ""
// 	for i := 0; i <= num-1; i++ {
// 		rom += "I"
// 	}
// 	return rom
// }

// func ConvertToRoman(arabic int) string {
// 	var result strings.Builder

// 	for arabic >= 1 {
// 		switch {
// 		case arabic >= 10:
// 			result.WriteString("X")
// 			arabic -= 10
// 		case arabic >= 9:
// 			result.WriteString("IX")
// 			arabic -= 9
// 		case arabic >= 5:
// 			result.WriteString("V")
// 			arabic -= 5
// 		case arabic >= 4:
// 			result.WriteString("IV")
// 			arabic -= 4
// 		default:
// 			result.WriteString("I")
// 			arabic--
// 		}
// 	}
// 	return result.String()
// }

// Refactor to decouple data and behaviors
type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
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
