package main

import "fmt"

func main(){
	// names := [...]string{"Alce", "Bob","Yoga","Chen","Eko","Dani"}
	// slice1 := names[4:6]

	// fmt.Println(slice1)

	// slice2 := names[:3]
	// fmt.Println(slice2)

	// slice3 := names[3:]
	// fmt.Println(slice3)

	// // Cara manual nya 
	// var slice4 [] string = names[:]
	// fmt.Println(slice4)

	// Function Slice

	// Append
	// days := [...]string {"Senin","Selasa","Rabu","Kamis","Jumat","Sabtu","Minggu"}
	// daySlice1 := days[5:]
	// daySlice1[0] = "Sabtu Baru"
	// daySlice1[1] = "Minggu Baru"
	// fmt.Println(days)

	// daySlice2 := append(daySlice1, "Libur Baru")
	// daySlice2[0] = "Ups"
	// fmt.Println(daySlice2)
	// fmt.Println(days)

	// Make Slice
	// var newSlice[] string = make([]string ,2,5)
	// newSlice[0]= "Yoga"
	// newSlice[1]= "Eka"
	// // newSlice[2]= "Pratama" harus append 

	// fmt.Println(newSlice)
	// fmt.Println(len(newSlice))
	// fmt.Println(cap(newSlice))

	// newSlice2 := append(newSlice, "Pratama")
	// fmt.Println(newSlice2)
	// fmt.Println(len(newSlice2))
	// fmt.Println(cap(newSlice2))

	// newSlice2[0] = "Budi"
	// fmt.Println(newSlice2)
	// fmt.Println(newSlice)

	// Copy Slice
	// fromSlice := days[:]
	// toSlice := make([]string, len(fromSlice), cap(fromSlice))

	// copy(toSlice, fromSlice)
	// fmt.Println(fromSlice)
	// fmt.Println(toSlice)


	// INI ARRAY 
	iniArray := [...]int {1,2,3,4,5}
	fmt.Println(iniArray)
	// INI SLICE
	iniSlice := []int {1,2,3,4,5}
	fmt.Println(iniSlice)

	fromSlice := iniSlice[:3]
	fmt.Println(fromSlice)

}	