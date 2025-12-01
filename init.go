package main

import (
	"golang-dasar/database"
	_ "golang-dasar/internal" // agar tidak error agar bisa tetep panggil init
	"fmt"
)

func main(){
	fmt.Println(database.GetDatabase())
}