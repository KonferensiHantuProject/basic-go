package main

import "fmt"

func main() {

	// Membuat array (dalam go array punya length yang fixed)
	// var ages [3]int = [3]int{20, 25, 30}
	var agesSecond = [3]int{20, 25, 30}

	names := [4]string{"Boni", "Mario", "Bobi", "Bobon"}
	names[1] = "Hartanto"

	fmt.Println(agesSecond, len(agesSecond))
	fmt.Println(names, len(names))

	// Slice bisa di manipulasi
	var scores = []int{31, 45, 150}
	scores[2] = 12

	// Menggunakan append untuk menambah isi slice
	scores = append(scores, 66)
	fmt.Println(scores)

	// Slice ranges (yang kedua tidak akan diambil)
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Dodi")
	fmt.Println(rangeOne)
}
