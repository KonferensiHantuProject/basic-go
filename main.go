package main

import "fmt"

func main() {
	name := "Bone"
	age := 41

	// Print
	fmt.Print("Halo, \n")
	fmt.Print("Dunia \n")
	fmt.Print("Baris Baru")

	// Otomatis line baru (Println)
	fmt.Println("Halo Bone")
	fmt.Println("Dadah Bone")
	fmt.Println("Umurku adalah", age, "dan namaku adalah", name)

	// Foramatting string (Printf) %_ = format specifier
	fmt.Printf("Umurku adalah %v dan namaku adalah %v \n", age, name)
	fmt.Printf("Umurku adalah %q dan namaku adalah %q \n", age, name)
	fmt.Printf("Umurku bertipe %T \n", age)
	fmt.Printf("Kamu mendapat skor %f poin \n", 55.5)
	fmt.Printf("Kamu mendapat skor %0.2f poin \n", 55.5) // Membatasi jumlah angka di belakang koma

	// Sprintf (Menyimpan string yang sudah di format)
	var str = fmt.Sprintf("Umurku adalah %v dan namaku adalah %v \n", age, name)
	fmt.Println("String yang di simpan adalah", str)
}
