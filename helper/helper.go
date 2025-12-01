package helper

// Akses modifier
var version = "1.0.0" // tidak bisa diakses diluar package karena awal nya huruf kecil
var Application = "golang" // bisa diakses di luar package

// Tidak bisa diakses diluar package
func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string{
	return "Hello " + name
}


func Contoh(){
	sayGoodBye("Eko")
	fmt(sayGoodBye)
}