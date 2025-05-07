package main

import "fmt"

type path []byte

// Aún pasandole una copia del slice, los cambios se reflejan en el original porque estamos modificando el array, y el slice en su estructura
// interna utiliza un puntero a un array.
func (p *path) ToUpper() {
	for i, b := range *p {
		if 'a' <= b && b <= 'z' {
			(*p)[i] = b + 'A' - 'a'
		}
	}
}

func main() {
	pathName := path("/usr/bin/t你o")
	pathName.ToUpper()
	fmt.Printf("%s\n", pathName)
	InitProblem()

	prueba := "/hola/que"
	fmt.Println(prueba)

	var arr5 [5]int = [5]int{1, 2, 3, 4, 5} //longitud 5, capacidad 5
	slice := arr5[:]
	slice2 := make([]int, len(slice), cap(slice)+1)
	copy(slice2, slice)
	fmt.Println("slice", slice)

	slice2 = slice2[:len(slice2)+1]
	slice2[5] = 6
	fmt.Println("slice2", slice2)
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	sliceString := make([]byte, len(sample), len(sample))
	copy(sliceString, sample)

	fmt.Println("Println string:")
	fmt.Println(sample)

	fmt.Println("Println slice:")
	fmt.Println(sliceString)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)

	myString := "Hey! ñaño!"
	fmt.Printf("String: %+q\n", myString)

	const nihongo = "日本語"

	for i, rune := range nihongo {
		fmt.Printf("Unicode: %#U at index: %d\n", rune, i)
	}

}
