package iteration

import "testing"
import "fmt"

func TestIteration(t *testing.T){
	t.Run("iterate a 5 times", func(t *testing.T){
		got := repeat("a", 5)
		want := "aaaaa"
	
		if got != want{
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("iterate b 3 times", func(t *testing.T){
		got := repeat("b",3)
		want := "bbb"

		if got != want{
			t.Errorf("got %q but want %q", got, want)
		}
	})
}


func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a", 5)
	}
}

func ExampleRepeat(){
	repetition := repeat("b", 5)
	fmt.Println(repetition)

	// Output: bbbbb
}