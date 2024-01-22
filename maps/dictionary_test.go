package maps

import "testing"


func TestSearch(t *testing.T) {
	t.Run("Known Word", func(t *testing.T) {
		dictionary := Dictionary{"test": "Ram is a tst subject in nepali vyakaran"}

		got,_ := dictionary.Search("test")
		want := "Ram is a tst subject in nepali vyakaran"
		
		assertString(t, got, want)
	})

	t.Run("Unknown Word", func(t *testing.T) {
		dictionary := Dictionary{"test": "Ram is a tst subject in nepali vyakaran"}

		_,got := dictionary.Search("somthing else")
		want := ErrNotFound

		assertError(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("New word", func(t *testing.T) {
		word := "test"
		definition := "Ram is a test subject in Nepali vyakaran"

		dictionary := Dictionary{}
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Existing word", func(t *testing.T) {

		word := "test"
		definition := "Ram is a tst subject in Nepali vyakaran"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "Ram is a test subject in Nepali vyakaran"

		dictionary := Dictionary{word: definition}
		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("New word", func(t *testing.T) {
		word := "test"
		definition := "Ram is a test subject in Nepali vyakaran"

		dictinoary := Dictionary{}
		err := dictinoary.Update(word, definition)

		assertError(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Deleting word", func(t *testing.T) {
		word := "test"
		definition := "Ram is a test subject in Nepali vyakaran"

		dictionary := Dictionary{word: definition}

		dictionary. Delete(word)

		_, err := dictionary.Search(word) 

		if err == nil {
			t.Errorf("Expeced %q to be deleted", word)
		}

	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, d Dictionary, word, definition string) {
	t.Helper()
	_, err := d.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}
}