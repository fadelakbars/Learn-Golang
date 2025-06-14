package main

import "fmt"

func main()  {
	sapa(namanya, "Fika")
}

// FUNGSI YANG MENERIMA CALLBACK
func sapa(callback func(string), nama string)  {
	callback(nama)
}

// FUNGSI CALLBACK
func namanya(nama string) {
	fmt.Println("Hai ", nama)
}

