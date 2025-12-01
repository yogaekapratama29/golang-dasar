package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func main(){
	// var nama Customer
	// nama.Name = "Yoga Eka Pratama"
	// nama.Address = "Purwokerto"
	// nama.Age = 20

	// fmt.Println(nama)

	// Stuct Literals

	// nama := Customer{
	// 	Name: "Yoga",
	// 	Address :  "Tegal",
	// 	Age : 21,
	// }

	// fmt.Println(nama)

	budi:= Customer {"Yoga","Slawi",22}
	fmt.Println(budi)
}

