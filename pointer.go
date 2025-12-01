package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main(){
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 //copy value
	address2 := &address1 // pointer

	address2.City = "Bandung"

	// astrisk

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1) // tidak berubah
	fmt.Println(address2) // berubah menjadi Bandung
}