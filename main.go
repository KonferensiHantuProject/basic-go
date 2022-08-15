package main

import (
	"fmt"
)

func main() {
	// x := 0

	// // Looping While
	// for x < 5 {
	// 	fmt.Println("Nilai x adalah:", x)
	// 	x++
	// }

	// Traditional For Loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Nilai i adalah:", i)
	// }

	names := []string{"mario", "luigi", "toni", "chopper", "drag"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// loooping pada variabel names (getting the index and the value)
	for index, value := range names {
		fmt.Printf("Posisinya ada pada index %v dengan nilainya adalah %v \n", index, value)
	}

	// Jika tidak ingin menggunakan index atau value bisa diganti dengan _
	for _, value := range names {
		fmt.Printf("Nilainya adalah %v \n", value)
		value = "string baru"
	}

	// Value tidak akan ke update (originalnya) karena itu seperti cuma variabel lokal di dalam looping itu saja
	fmt.Println(names)
}
