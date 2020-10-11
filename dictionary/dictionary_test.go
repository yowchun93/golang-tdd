package dictionary

import "testing"

// func TestSearch(t *testing.T) {
// 	dictionary := map[string]string{"test": "this is just a test"}

// 	got := Search(dictionary, "test")
// 	want := "this is just a test"

// 	if got != want {
// 		t.Errorf("got %q want %q given, %q", got, want, "test")
// 	}
// }

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

// Dictionary as a custom type
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got, _ := dictionary.Search("test")
	want := "this is just a test"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, err := dictionary.Search("random")
		// want := "no such word"

		if err == nil {
			t.Fatal("expected to get an error.")
		}
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, _ := dictionary.Search("test")

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		if err == nil {
			t.Fatal("expected to get an error.")
		}
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	newDefinition := "new definition"

	dictionary.Update(word, newDefinition)

	assertDefinition(t, dictionary, "test", "new definition")

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		if err == nil {
			t.Fatal("expected to get an error.")
		}
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}
