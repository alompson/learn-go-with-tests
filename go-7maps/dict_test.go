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
	dict := Dict{}
	word := "paulistano"
	def := "ser estranho de dialetos desconhecidos"

	dict.Add(word, def)

	assertDefinition(t, dict, word, def)
}

func assertDefinition(t testing.TB, dict Dict, word, def string) {
	got, err := dict.Search(word)

	if err != nil{
		t.Fatal("Should find Paulistano:", err)
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