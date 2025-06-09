package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Se esperaba obtener %q, pero se ha obtenido %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for range b.N {
		Repeat("a")
	}
}
