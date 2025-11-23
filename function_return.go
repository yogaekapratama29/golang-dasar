package main

import "fmt"

func getHello(name string) string {
	return "Hello" + name
}

func main(){
	result := getHello("Yoga")
	fmt.Println(result)
}