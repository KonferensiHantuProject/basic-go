package main

import (
	"fmt"
	"sort"
)

func main() {

	// greetings := "Halo teman-teman semua"

	// fmt.Println(strings.Contains(greetings, "Halo"))          // Mencari kata tertentu dalam kalimat
	// fmt.Println(strings.ReplaceAll(greetings, "Halo", "Hai")) // Mengembalikan string baru yang sudah diubah sesuai dengan yang digantikan
	// fmt.Println(strings.ToUpper(greetings))
	// fmt.Println(strings.Index(greetings, "t"))
	// fmt.Println(strings.Split(greetings, " "))

	// // Original
	// fmt.Println("Yang original adalah", greetings)

	ages := []int{45, 20, 21, 65, 87, 23, 44, 56, 67, 30}

	sort.Ints(ages)
	fmt.Println(ages)

	// Searching the integers
	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoyo", "nono", "momo", "dragon", "ball"}

	sort.Strings(names)
	fmt.Println(names)

	// Searching the slices (strings)
	fmt.Println(sort.SearchStrings(names, "ball"))
}
