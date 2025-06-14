package main

import "fmt"

func main()  {
	nilaiLama := 1
	fmt.Println("Nilai awal diluar fungsi : ", nilaiLama)
	
	ubahnilaipointer(&nilaiLama, 2)
	fmt.Println("Nilai baru diluar fungsi : ", nilaiLama)
}

func ubahnilaipointer(nilaiLamaPointer *int, nilaiBaruPointer int)  {
	*nilaiLamaPointer = nilaiBaruPointer
	fmt.Println("Nilai baru dalam fungsi biasa = ", *nilaiLamaPointer)
}