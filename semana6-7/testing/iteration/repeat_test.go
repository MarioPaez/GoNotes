package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Call Repeat with repeatedCount equals to 10", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("Se esperaba obtener %q, pero se ha obtenido %q", expected, repeated)
		}
	})

	t.Run("Call Repeat with repeatedCount equals to 3", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		if repeated != expected {
			t.Errorf("Se esperaba obtener %q, pero se ha obtenido %q", expected, repeated)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	Repeat("a", b.N)
}

func ExampleRepeat() {
	got := Repeat("m", 3)
	fmt.Println(got)
	//Output: mmm
}
