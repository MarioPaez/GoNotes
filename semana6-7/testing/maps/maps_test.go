package maps

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("get a word that exist", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("word does not exist", func(t *testing.T) {
		_, got := dictionary.Search("noExist")

		if got == nil {
			t.Fatal("A error must to be throw")
		}
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	t.Run("adding new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("word already exist", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		got := dictionary.Add(word, definition)
		if got == nil {
			t.Fatal("A error must to be throw")
		}

		assertError(t, got, ErrWordExist)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		newDefinition := "new definition"

		err = dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q but want %q", got, want)
	}
}

func assertError(t testing.TB, got, err error) {
	t.Helper()
	if got != err {
		t.Errorf("Got %q but want %q", got, err)
	}
}
