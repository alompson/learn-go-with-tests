package main

import "testing"
import "reflect"

func TestSum(t *testing.T){

	t.Run("Array of size 5", func(t *testing.T){
		numbers :=[]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
	
		if got != want{
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	// t.Run("Array of any size", func(t *testing.T){
	// 	numbers :=[]int{1, 2, 3}

	// 	got := Sum(numbers)
	// 	want := 6
	
	// 	if got != want{
	// 		t.Errorf("got %d want %d given %v", got, want, numbers)
	// 	}
	// })
}

func TestSumAll(t *testing.T){
	t.Run("with 2 arrays", func(t *testing.T){

		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T){

	checkSums := func(t testing.TB, got, want []int){
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	}

	t.Run("with 2 arrays", func(t *testing.T){

		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("with empty array", func(t *testing.T){

		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}