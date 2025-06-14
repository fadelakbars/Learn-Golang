package main

import "fmt"

/*
	struct adalah tipe data buatan untuk menampung beberapa data berbeeda dalam satu kesatuan
*/

type Mahasiswa struct {
	Nama string
	Umur int
	Email string
}

func main()  {
	var mhs1 Mahasiswa
	mhs1.Nama = "Fadhel"
	mhs1.Umur = 21
	mhs1.Email = "fadhelakbarsallang@gmail.com"
	fmt.Println(mhs1)

	// STRUCT LITERAL
	mhs2 := Mahasiswa{
		Nama: "Akbar",
		Umur: 29,
		Email: "akabarnurinsan@gmail.com",
	}
	fmt.Println(mhs2)
}