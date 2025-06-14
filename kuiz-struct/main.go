package main

import "fmt"

type Prodi struct {
	NamaProdi string
	MataKuliah []string
}

func (prodi *Prodi) tambahMatkul(matkulBaru string) {
	prodi.MataKuliah = append(prodi.MataKuliah, matkulBaru)
}

func main()  {
	matkul := []string{"Alprogdas", "PIK"}
	prodi := Prodi{NamaProdi: "PTIK", MataKuliah: matkul}

	prodi.tambahMatkul("Web Programing")

	for _, namaMatkul := range prodi.MataKuliah {
		fmt.Println(namaMatkul)
	}
}