package main

import "fmt"

func main() {
	x := 2

	var pointer *int = &x

	fmt.Println("isi x : ",x)
	fmt.Println("alamat x : ",&x)
	fmt.Println("alamat pointer : ", pointer)
	// DEREFERNCING UNTUK MENCETAK VALUE POINTER
	fmt.Println("isi pointer : ", *pointer)

	y := 3
	fmt.Println("nilai awal y : ", y)
	nilaibaruy := &y
	*nilaibaruy = 3
	fmt.Println("nilai baru y : ", *nilaibaruy)

}
