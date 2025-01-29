package main

import "testing"

func TestSearch(t *testing.T) {
	d := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := d.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := d.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := d.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, d, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word, definition := "bybis", "kiausai"
		d := Dictionary{word: definition}

		err := d.Add(word, "tava tevs")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, d, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dict.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("unknown word", func(t *testing.T) {
		word := "test"
		newDefinition := "new definition"
		dict := Dictionary{}

		err := dict.Update(word, newDefinition)

		assertError(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{word: "test def"}

		dict.Delete(word)

		_, err := dict.Search(word)
		assertError(t, err, ErrNotFound)
	})

	t.Run("unknown word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}

		err := dict.Delete(word)

		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got: %q; want: %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got: %q; want: %q", got, want)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}
