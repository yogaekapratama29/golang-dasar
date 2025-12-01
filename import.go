package main


import (
	"golang-dasar/helper"
	"fmt"
)


func main(){
	result := helper.SayHello("Yoga")
	fmt.Println(result)

	fmt.Println(helper.Application)
	fmt.Println(helper.version) // tidak bisa diakses
	fmt.Println(helper.sayGoodBye("Eko")) // tidak bisa diakses
}