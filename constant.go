package main 

import "fmt"

func main() {
	// const firstName = "Yoga"
	// const age = 20

	// firstName = "Carmen" // akan error karena constant tidak bisa diubah nilainya

	// fmt.Println(firstName)
	// fmt.Println(age)

	// Multiple Constant
const(
	firstName = "Carmen"
	lastName = "Karina"
)

fmt.Println(firstName)
fmt.Println(lastName)
}