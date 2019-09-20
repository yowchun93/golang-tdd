package roman

import "testing"

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

func TestRomanNumerals(t *testing.T) {
	// Struct literals
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to I", 2, "II"},
		{"3 gets converted to I", 3, "III"},
		{"4 gets converted to IV", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
		{"8 gets converted to IX", 8, "VIII"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"40 gets converted to XL", 40, "XL"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Want

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestArabicNumerals(t *testing.T) {
	got := ConvertToArabic("I")
	want := 1

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	t.Run("Test II returns 2", func(t *testing.T) {
		got := ConvertToArabic("II")
		want := 2

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Test III returns 3", func(t *testing.T) {
		got := ConvertToArabic("III")
		want := 3

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Test IV returns 4", func(t *testing.T) {
		got := ConvertToArabic("IV")
		want := 4

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
