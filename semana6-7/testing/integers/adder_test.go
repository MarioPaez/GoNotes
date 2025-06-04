package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("Test suma", func(t *testing.T) {
		got := Add(2, 3)
		want := 5

		if got != want {
			t.Errorf("Se esperaba un '%d' y se ha obtenido un '%d'", want, got)
		}

	})
}
