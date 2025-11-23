package main

import "fmt"

func main(){

	// name := "Eka"

	// if name == "Yoga" {
	// 	fmt.Println("Hello Yoga")
	// }else if name == "Eka"{
	// 	fmt.Println("Hello Eka")
	// }else {
	// 	fmt.Println("Hi, boleh kenalan?")
	// }

	name := "Yoga"

	// if short statement 
	if length := len(name); length > 4 {
		fmt.Println("Nama terlalu panjeng")
	}else {
		fmt.Println("Nama sudah benar")
	}
}