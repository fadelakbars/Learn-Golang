package main

import "fmt"

func main()  {
	nomor := 1
	// MEMBUAT VARIABEL SEBAGIA POINTER YANG MEREFERENCES KE ALAMAT MEMORI YANG SAMA DENGAN VARIABEL nomor
	nomorB := &nomor

	// OUTPUT NORMAL, MENCETAK VALUE variabel nomor
	fmt.Println(nomor)

	// AKAN MENCETAK ALAMAT MEMORI VARIABEL nomor
	fmt.Println(nomorB)

	// DEREFERENCES DENGAN * UNTUK MENCETAK VALUE PADA ALAMAT MEMORI TERSEBUT
	fmt.Println(*nomorB)

	// MENGUBAH VALUE DENGAN DENGAN POINTER
	*nomorB = 10

	fmt.Println(nomorB)
	fmt.Println(nomor)
	
	// DEREFERENCES UNTUK MENCETAK VALUE POINTER
	fmt.Println(*nomorB)
}