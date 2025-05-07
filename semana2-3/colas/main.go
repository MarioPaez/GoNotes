package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type Cola struct {
	len  int
	head *Node
}

type ColaSlice struct {
	cola []int
}

func (c *Cola) Size() int {
	return c.len
}

func (c *Cola) IsEmpty() bool {
	return c.len == 0
}

func (c *Cola) Enque(val int) {
	newNode := &Node{val: val, next: nil}
	if c.IsEmpty() {
		c.head = newNode
		c.len++
		return
	}

	current := c.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
	c.len++
}

func (c *Cola) Deque() (int, error) {
	if c.IsEmpty() {
		return 0, fmt.Errorf("no se puede eliminar porque la cola está vacía")
	}

	valReturn := c.head.val
	c.head = c.head.next
	c.len--
	return valReturn, nil
}

func (c *Cola) Peek() (int, error) {
	if c.IsEmpty() {
		return 0, fmt.Errorf("no hay elementos en la cola")
	}

	return c.head.val, nil
}

func (c *Cola) Print() {
	if c.IsEmpty() {
		fmt.Println("La cola está vacía")
		return
	}

	fmt.Println("Contenido de la cola:")
	current := c.head
	pos := 0
	for current != nil {
		fmt.Printf("Posición %d: %d\n", pos, current.val)
		current = current.next
		pos++
	}
}

func (c *ColaSlice) Size() int {
	return len(c.cola)
}

func (c *ColaSlice) IsEmpty() bool {
	return len(c.cola) == 0
}

func (c *ColaSlice) Enque(val int) {
	c.cola = append(c.cola, val)
}

func (c *ColaSlice) Deque() (int, error) {
	if c.IsEmpty() {
		return 0, fmt.Errorf("no se puede eliminar porque la cola está vacía")
	}
	res := c.cola[0]
	c.cola = c.cola[1:]
	return res, nil
}

func (c *ColaSlice) Peek() (int, error) {
	if c.IsEmpty() {
		return 0, fmt.Errorf("no hay elementos en la cola")
	}
	val := c.cola[0]
	return val, nil
}

func (c *ColaSlice) Print() {
	if c.IsEmpty() {
		fmt.Println("La cola está vacía")
		return
	}

	fmt.Println("Contenido de la cola:")

	for i, val := range c.cola {
		fmt.Printf("Posición %d: %d\n", i, val)
	}
}

func main() {
	// c := &Cola{}

	// fmt.Println("Cola creada.")
	// fmt.Printf("¿Está vacía?: %v\n", c.IsEmpty())
	// fmt.Printf("Tamaño: %d\n", c.Size())

	// fmt.Println("\nEncolando elementos 10, 20, 30...")
	// c.Enque(10)
	// c.Enque(20)
	// c.Enque(30)

	// c.Print()
	// fmt.Printf("¿Está vacía?: %v\n", c.IsEmpty())
	// fmt.Printf("Tamaño: %d\n", c.Size())

	// fmt.Println("\nViendo el primer elemento con Peek:")
	// val, err := c.Peek()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Valor al frente:", val)
	// }

	// fmt.Println("\nDesencolando elementos:")
	// for !c.IsEmpty() {
	// 	val, err := c.Deque()
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 	} else {
	// 		fmt.Println("Elemento desencolado:", val)
	// 	}
	// 	c.Print()
	// }

	// fmt.Printf("¿Está vacía ahora?: %v\n", c.IsEmpty())
	// fmt.Printf("Tamaño final: %d\n", c.Size())

	// fmt.Println("\nIntentando hacer Peek y Deque con cola vacía:")
	// _, err = c.Peek()
	// if err != nil {
	// 	fmt.Println("Peek error:", err)
	// }

	// _, err = c.Deque()
	// if err != nil {
	// 	fmt.Println("Deque error:", err)
	// }

	c1 := &ColaSlice{}

	fmt.Println("ColaSlice creada.")
	fmt.Printf("¿Está vacía?: %v\n", c1.IsEmpty())
	fmt.Printf("Tamaño: %d\n", c1.Size())

	fmt.Println("\nEncolando elementos 100, 200, 300...")
	c1.Enque(100)
	c1.Enque(200)
	c1.Enque(300)

	c1.Print()
	fmt.Printf("¿Está vacía?: %v\n", c1.IsEmpty())
	fmt.Printf("Tamaño: %d\n", c1.Size())

	fmt.Println("\nUsando Peek: ")
	val1, err := c1.Peek()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Primer elemento:", val1)
	}

	fmt.Println("\nDesencolando todos los elementos:")
	for !c1.IsEmpty() {
		val, err := c1.Deque()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Elemento eliminado:", val)
		}
		c1.Print()
	}

	fmt.Println("\nIntentando Peek y Deque con cola vacía:")
	_, err = c1.Peek()
	if err != nil {
		fmt.Println("Peek error:", err)
	}

	_, err = c1.Deque()
	if err != nil {
		fmt.Println("Deque error:", err)
	}
}
