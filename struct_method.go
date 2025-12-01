package main

import "fmt"

type Customer struct{
	Name string
	Age int
}

func(customer Customer)sayHello(name string){
	fmt.Println("Hello,",name, "My Name is", customer.Name)
}

func main(){
	yoga := Customer{"Yoga",20}
	yoga.sayHello("Eko")
	fmt.Println(yoga)
}