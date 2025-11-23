package main


import "fmt"

func main(){
	// counter := 1


	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	// fmt.Println("Selesai")

	// For dengan statement
	// init statement, statement sebelum for dieksekusi
		// post statement, statement yang akan dieksekusi diakhir tiap perulangan
	for counter:= 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// For Range
	// For bisa digunakan untuk melakukan iterasi terhadap semua data collection
	// Data collection contoh nya Array, Slice dan Map

	names := []string{"Yoga","Eka","Pratama"}
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	// Manual
	// names := []string{"Eko","Kurniawan","Yoga"}
	// for i:= 0; i < len(names);i++{
	// 	fmt.Println(names[i])
	// }


}