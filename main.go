package main

import "fmt"

func main() {

	// Strings
	var nameOne string = "Nama"
	var nameTwo = "Luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	// Updating Value of variable
	nameOne = "Nama Update"
	nameThree = "Update"

	fmt.Println(nameOne, nameTwo, nameThree)

	// Deklarasi variabel di awal dan tidak bisa digunakan diluar function
	nameFour := "yosha"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// Numbers (Integer and decimal)
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// Bits dan memories (uint tidak menerima negatif)
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint16 = 2556

	// Float (memiliki desimal)
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 893821948298498234983948398.9
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
