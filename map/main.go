package main

import "fmt"

func main() {

	hash := map[int]string{}

	hash[0] = "Mario"
	hash[1] = "Luigi"
	hash[2] = "Peach"
	hash[3] = "Yoshi"
	hash[4] = "Toad"
	hash[5] = "Bowser"
	hash[6] = "Wario"
	hash[7] = "Donkey Kong"
	hash[8] = "Link"
	hash[9] = "Zelda"
	hash[10] = "Kirby"
	hash[11] = "Samus"

	fmt.Println(hash)

	m := make(map[string]map[string]int)

	mm, ok := m["Mario"]

	if !ok {
		mm = make(map[string]int)
		m["Mario"] = mm
	}

	mm["Luigi"]++

}
