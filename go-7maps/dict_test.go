package dictionary

import(
	"testing"
)

func TestSearch(t *testing.T){
	t.Run("Known Word", func(t *testing.T){
		dict := Dict{"test": "this is a test"}

		got, _ := dict.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("Unknown Word", func(t *testing.T){
		dict := Dict{}

		_, got := dict.Search("test")

		if got == nil{
			t.Fatal("Should get an error")
		}

		assertError(t, got, ErrNotFound)
		//assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T){
	t.Run("New word", func (t *testing.T){
		dict := Dict{}
		word := "paulistano"
		def := "ser estranho de dialetos desconhecidos"
	
		err := dict.Add(word, def)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, def)

	})

	t.Run("Existing word", func (t *testing.T){
		word := "test"
		def := "this is just a test"
		dict := Dict{word: def}

		err := dict.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, def)

	})

}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T){
		word := "test"
		def := "this is a test"
		dict := Dict{word: def}
	
		newDef := "a new test!"
	
		err := dict.Update(word, newDef)
		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDef)
	})

	t.Run("new word", func(t *testing.T){
		dict := Dict{}
		word := "test"
		def := "this is a test"

		err := dict.Update(word, def)
		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T){
	t.Run("existing word", func(t *testing.T){
		word := "test"
		def := "this is a test"
		dict := Dict{word: def}

		dict.Delete(word)
		_, err := dict.Search(word)
		assertError(t, err, ErrNotFound)
	})

	t.Run("new word", func(t *testing.T){
		dict := Dict{}
		word := "test"

		err := dict.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dict Dict, word, def string) {
	t.Helper()
	got, err := dict.Search(word)

	if err != nil{
		t.Fatal("Should find word:", err)
	}

	assertStrings(t, got, def)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error){
	t.Helper()

	if got != want{
		t.Errorf("got error %q want %q", got, want)
	}
}

