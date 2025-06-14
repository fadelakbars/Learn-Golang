package main

import "fmt"

type Mahasiswa struct{
	Nama string
}

func ubahNama(atribut *Mahasiswa, namaBaru string)  {
	atribut.Nama = namaBaru
}

func main()  {
	mhs1 := Mahasiswa{Nama: "Fadhel"}
	fmt.Println(mhs1.Nama)

	ubahNama(&mhs1, "Fadhel Akbar")
	fmt.Println(mhs1.Nama)
}