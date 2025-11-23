package main

import "fmt"

func getComplateName() (firstName, middleName, lastName string) {
	firstName = "Yoga"
	middleName = "Eka"
	lastName = "Pratama"

	return firstName, middleName, lastName
}

func main(){
	firstName, middleName, lastName := getComplateName()
	fmt.Println(firstName, middleName, lastName)

}