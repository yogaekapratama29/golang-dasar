package main

import "fmt"

func main(){ 

	//  Bisa manual dari awal
	// var person map[string]string = map[string]string{}
	// person["name"] = "Yoga"
	// person["address"] = "Indonesia"
	// fmt.Println(person)

	// Lebih singkat 
	// person := map[string]string{
	// 	"name" : "Yoga",
	// 	"address" : "Indonesia",
	// }

	// fmt.Println(person["name"])
	// fmt.Println(person["address"])
	// fmt.Println(person) 

	// Function Map

	book:= make(map[string]string)

	book["title"] = "Buku Golang"
	book["author"] = "Yoga Pratama"
	book["wrong"] = "Ups"

	delete(book,"wrong")

	fmt.Println(book)


}