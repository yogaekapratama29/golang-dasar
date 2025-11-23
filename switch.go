package main


import "fmt"

func main(){
	name := "Yoga"

	// switch name {
	// case "Yoga" :
	// 	fmt.Println("Hello Yoga")
	// case "Budi" :
	// 	fmt.Println("Hello Budi")
	// default : 
	// 	fmt.Println("Hello User")
	// }


	// Switch with short statement
	// switch length := len(name); length > 5 {
	// case true :
	// 	fmt.Println("Nama terlalu panjang")
	// case false :
	// 	fmt.Println("Nama sudah benar")
	// }

	// Switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10 : 
	fmt.Println("Nama terlalu panjang")
	case length > 5 : 
	fmt.Println("Nama lumayan panjang")
	default : 
	fmt.Println("Nama sudah benar")
	}
}