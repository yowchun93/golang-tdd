package dictionary

import "testing"

// version 1
// func TestSearch(t *testing.T) {
// 	dictionary := map[string]string{"test": "this is just a test"}

// 	got := Search(dictionary, "test")
// 	want := "this is just a test"

// 	if got != want {
// 		t.Errorf("got %q want %q given, %q", got, want, "test")
// 	}
// }

// version 2
func TestSearch(t *testing.T) {
	// dictionary := map[string]string{"test": "this is just a test"}
	dictionary := Dictionary{"test": "this is just a test"}
	got, _ := dictionary.Search("test")
	want := "this is just a test"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func TestUnknownWord(t *testing.T) {
	// dictionary := map[string]string{"test": "this is just a test"}
	dictionary := Dictionary{"test": "this is just a test"}
	_, error := dictionary.Search("lol")

	if error == nil {
		t.Fatal("expected Error")
	}
}

func TestAddWord(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		dictionary.Add("lol", "wohoo")

		want := "wohoo"
		got, _ := dictionary.Search("lol")
		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "wohoo")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, "test", "this is just a test")
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "just a test la"
	dictionary := Dictionary{word: definition}
	newDefinition := "new definition"

	dictionary.Update(word, newDefinition)
	assertDefinition(t, dictionary, word, newDefinition)

	// declaring anonymous function
	t.Run("non-existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "test only"}
		err := dictionary.Update("lol", "hehe")

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

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
