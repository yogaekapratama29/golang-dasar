package main

import "fmt"


func main(){
	// var names [3] string
	// names[0] = "Yoga"
	// names[1] = "Eka"
	// names[2] = "Pratama"

	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[2])

	// var values = [3]int{
	// 	90,
	// 	80,
	// }
	// fmt.Println(values[0])
	// fmt.Println(values[1])
	// fmt.Println(values[2])

	var values = [...] int {
		90,
		70,
		50,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)


}