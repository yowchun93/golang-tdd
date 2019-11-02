package roman

import "testing"

func TestRomanNumeral(t *testing.T) {

	type test struct {
		description string
		arabic      int
		want        string
	}

	tests := []test{
		{"test 1 return I", 1, "I"},
		{"test 2 return II", 2, "II"},
		{"test 3 return III", 3, "III"},
		{"test 4 return IV", 4, "IV"},
		{"test 5 return V", 5, "V"},
		{"6 gets converted to VI", 6, "VI"},
		{"7 gets converted to VII", 7, "VII"},
		{"8 gets converted to VIII", 8, "VIII"},
		{"10 gets converted to X", 10, "X"},
	}

	for _, test := range tests {
		got := ToRoman(test.arabic)
		want := test.want

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
}

func TestToArabic(t *testing.T) {
	type test struct {
		description string
		roman       string
		want        int
	}

	tests := []test{
		{"test 1 return I", "I", 1},
		{"test 2 return II", "II", 2},
		{"test 3 return III", "III", 3},
		{"test 4 return IV", "IV", 4},
		{"test 5 return V", "V", 5},
		{"test 6 return VI", "VI", 6},
		{"10 gets converted to X", "X", 10},
	}

	for _, test := range tests {
		got := ToArabic(test.roman)
		want := test.want

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}
