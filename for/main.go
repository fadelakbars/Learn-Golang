package main

import "fmt"

func main()  {

/*
	STRUKTU DASAR FOR
	for inisialisasi; kondisi; post-statement {
		// kode yang akan diulang
	}

*/

	// REGULER FOR
	fmt.Println("FOR REGULER")
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}
	
	// FOR TANPA INISIALISASI
	fmt.Println("\nFOR TANPA INISIALISASI")
	i := 1
	for i <= 5 {
		fmt.Println("Perulangan ke-", i)
		i++
	}
	
	// MENGHENTIKAN FOR DENGAN BREAK
	fmt.Println("\nMENGHENTIKAN FOR DENGAN BREAK")
	for j := 1; j <= 10; j++ {
		if j == 6 {
			fmt.Println("Iterasi dihentikan di", j)
			break
		}
		fmt.Println("Perulangan ke-", j)
	}

	// MENGGUNAKAN CONTINUE UNTUK MELEWATI ITERASI
	fmt.Println("\nMENGGUNAKAN CONTINUE UNTUK MELEWATI ITERASI")
	for k := 1; k <= 10; k++ {
		if k%2 == 0 {
			fmt.Println("Angka genap, melewati iterasi", k)
			continue
		}
		fmt.Println("Angka ganjil:", k)
	}

	// FOR RANGE
	fmt.Println("\nMENGGUNAKAN FOR RANGE")
	angkaAngka := []int{1, 2, 3, 4, 5}
	for index, angka := range angkaAngka {
		fmt.Printf("Indeks %d: Nilai %d\n", index, angka)
	}

	// MENGGUNAKAN _ UNTUK MENGABAIKAN INDEKS
	fmt.Println("\nMENGGUNAKAN _ UNTUK MENGABAIKAN INDEKS")
	angkaAngkaLain := []int{10, 20, 30, 40, 50}
	for _, angka := range angkaAngkaLain {
		fmt.Println("Nilai:", angka)
	}

	// LOOPING MAP
	fmt.Println("\nLOOPING MAP")
	dataMap := map[string]int{
		"apel":  10,
		"jeruk": 20,
		"mangga": 30,
	}

	for buah, jumlah := range dataMap {
		fmt.Printf("Buah: %s, Jumlah: %d\n", buah, jumlah)
	}

	// NESTED FOR LOOP
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

}