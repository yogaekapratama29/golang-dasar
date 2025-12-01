package main

import "fmt"

// Tidak bisa digunakan karena tipe data nil hanya bisa di interface, pointer, map, slice, channel, dan function
// func Contoh(name string) string {
// 	if name == "" {
// 		return nil
// 	}else {
// 		return name
// 	}
// }

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}else{
		return map[string]string{
			"name" : name,
		}
	}
}

func main(){
	data := NewMap("")
	if data == nil {
		fmt.Println("Data Map masih kosong")
	}else {
		fmt.Println(data["name"])
	}
}