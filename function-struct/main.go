package main

import "fmt"

type Mahasiswa struct {
	Nama string
	Umur int
}

/*
	EMBEDDED STRUCT
	artinya, sebuah struct juga bisa menjadi property pada struct lain
*/
type Kelas struct {
	Mahasiswa []Mahasiswa //mahasiswa adalah properti struct bertipe slice dari struct mahasiswa
	Keti   Mahasiswa
}

func main()  {
	mhs1 := Mahasiswa{}
	mhs1.Nama = "Fadel"
	mhs1.Umur = 21
	mhs2 := Mahasiswa{Nama: "Fika", Umur: 21}
	mhs3 := Mahasiswa{Nama: "Alya", Umur: 21}

	daftarMhs := []Mahasiswa{mhs1, mhs2}
	daftarMhs2 := []Mahasiswa{mhs2, mhs3}

	kelas := Kelas{daftarMhs, mhs3}
	kelas2 := Kelas{daftarMhs2, mhs1}
	
	//  MENCETAK MAHASISWA DENGAN STRUCT SEBAGAI PARAMETER FUNCTION
	fmt.Println(tampilkanMahasiswa(mhs1))
	fmt.Println("")
	
	//  MENCETAK DETAIL KELAS DENGAN FUNCTION STRUCT SEBAGAI PARAMTER
	tampilkanKelas(kelas)
	fmt.Println("")

	// MENCETAK MAHASISWA MELALUI METHOD STRUCT
	fmt.Println(mhs2.tampilkanDetailMahasiswa())
	fmt.Println("")

	// MENCETAK DETAIL KELAS DENGAN METHOD STRUCT
	kelas2.tampilkanDetailKelas()
}

// STRUCT SEBAGAI PARAMTER FUNCTION
func tampilkanMahasiswa(mahasiswa Mahasiswa) string {
	result := fmt.Sprintf("Nama : %s , Umur : %d", mahasiswa.Nama, mahasiswa.Umur)
	return result
}

// FUNGSI MENCETAK DETAIL KELAS DENGAN FUNGSI YANG MENERIMA PARAMTER STRUCT
func tampilkanKelas(kelas Kelas)  {
	fmt.Printf("Nama Keti : %s", kelas.Keti.Nama)
	fmt.Println("")
	fmt.Printf("Jumlah Mahasiswa : %d", len(kelas.Mahasiswa))
	
	fmt.Println("")
	for _, mahasiswa := range kelas.Mahasiswa {
		fmt.Println(mahasiswa.Nama)
	}
}

/* 
	METHOD STRUCT, FUNGSI KHUSUS UNTUK STRUCT
	method bisa dipanggil setelah object struct diinistansiasi
*/
func (mahasiswa Mahasiswa) tampilkanDetailMahasiswa() string {
	return fmt.Sprintf("Nama : %s , Umur : %d", mahasiswa.Nama, mahasiswa.Umur)
}

func (kelas Kelas) tampilkanDetailKelas() {
	fmt.Printf("Nama Keti : %s : ", kelas.Keti.Nama)
	fmt.Println("")
	
	for _, mahasiswa := range kelas.Mahasiswa {
		fmt.Println(mahasiswa.Nama)
	}
}