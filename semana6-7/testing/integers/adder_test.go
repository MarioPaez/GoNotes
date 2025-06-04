package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Test suma", func(t *testing.T) {
		got := Add(2, 3)
		want := 5

		if got != want {
			t.Errorf("Se esperaba un '%d' y se ha obtenido un '%d'", want, got)
		}

	})
}

func ExampleAdd() {
	sum := Add(9, 7)
	fmt.Println(sum)
	//Output: 16
}

func ExampleAdd_second() {
	sum := Add(9, 7)
	sum2 := Add(sum, 10)
	fmt.Println(sum2)
}
