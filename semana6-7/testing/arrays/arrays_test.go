package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestAll(t *testing.T) {
	t.Run("SumAll with 3 slices as input", func(t *testing.T) {
		arr1 := []int{1, 2, 3, 4, 5}
		arr2 := []int{1, 2, 3, 4}
		arr3 := []int{1, 2, 3, 4, 5, 6}

		got := SumAll(arr1, arr2, arr3)
		want := []int{15, 10, 21}

		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d given, 1: %v, 2: %v, 3: %v", got, want, arr1, arr2, arr3)
		}
	})

	t.Run("SumAll with nil as input", func(t *testing.T) {

		got := SumAll(nil)
		want := []int{0}
		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d given %v", got, want, nil)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, want, got []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 19}
		checkSums(t, got, want)

	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}
