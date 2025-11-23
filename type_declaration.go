package main

import "fmt"

func main(){

	type NoKTP string

	var noKTPYoga NoKTP = "321213231"

	var contoh string = "2222222"

	fmt.Println(noKTPYoga)

	var contohKTP NoKTP = NoKTP(contoh)
	fmt.Println(contohKTP)
}