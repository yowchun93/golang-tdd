package roman

import (
	"fmt"
	"testing"
)

// func TestRomanNumerals(t *testing.T) {
// 	got := ConvertToRoman(1)
// 	want := "I"

// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}

// 	t.Run("2 gets converted to II", func(t *testing.T) {
// 		got := ConvertToRoman(2)
// 		want := "II"

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})
// }

var cases = []struct {
	Description string
	Arabic      int
	Roman       string
}{
	{"1 gets converted to I", 1, "I"},
	{"2 gets converted to II", 2, "II"},
	{"3 gets converted to III", 3, "III"},
	{"4 gets converted to IV", 4, "IV"},
	{"5 gets converted to V", 5, "V"},
	{"6 gets converted to VI", 6, "VI"},
	{"9 gets converted to VI", 9, "IX"},
	{"10 gets converted to X", 10, "X"},
	{"14 gets converted to XIV", 14, "XIV"},
	{"18 gets converted to XVIII", 18, "XVIII"},
	{"20 gets converted to XX", 20, "XX"},
	{"39 gets converted to XXXIX", 39, "XXXIX"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

var newCases = []struct {
	Description string
	Arabic      int
	Roman       string
}{
	{"1 gets converted to I", 1, "I"},
	{"2 gets converted to II", 2, "II"},
	{"4 gets converted to IV", 4, "IV"},
	{"5 gets converted to V", 5, "V"},
	{"9 gets converted to VI", 9, "IX"},
	{"10 gets converted to X", 10, "X"},
	{"14 gets converted to XIV", 14, "XIV"},
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range newCases[:7] {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

// func TestRomanNumeralValues(t *testing.T) {
// 	want := 5
// 	got := allRomanNumerals.ValueOf("V")

// 	if got != want {
// 		t.Errorf("got %d, want %d", got, want)
// 	}

// 	t.Run("test IV", func(t *testing.T) {
// 		want := 4
// 		got := allRomanNumerals.ValueOf("IV")

// 		if got != want {
// 			t.Errorf("got %d, want %d", got, want)
// 		}
// 	})
// }
