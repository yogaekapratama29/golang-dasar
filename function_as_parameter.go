package main

import "fmt"

// func sayHelloWithFilter(name string, filter func(string) string){
// 	filterdNamed := filter(name)
// 	fmt.Println("Hello", filterdNamed)
// }

// Pakai Function Type Declaration untuk alias agar sedikit 
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "...."
	}else{
		return name
	}
}

func main(){
	sayHelloWithFilter("Yoga", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("anjing", filter)
}