package main

import "fmt"

type Mahasiswa struct {
	Nama string
}

func Yudisium(mahasiswa *Mahasiswa)  {
	mahasiswa.Nama = mahasiswa.Nama + " S.Pd"
}

func (mahasiswa *Mahasiswa) yudisiumMethod() {
	mahasiswa.Nama = mahasiswa.Nama + " S.Pd, M.Eng"
}

func main()  {
	fadel := Mahasiswa{Nama: "Fadhel Akbar"}

	Yudisium(&fadel)
	fmt.Println(fadel.Nama)
	
	fadel.yudisiumMethod()
	fmt.Println(fadel.Nama)
}