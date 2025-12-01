package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married(){
	man.Name = "Mr. " + man.Name
}

func main(){
	yoga := Man{"Yoga"}
	yoga.Married()

	fmt.Println(yoga.Name)
}