package main


import (

	"fmt"

	"kuiz-satu/kalkulasi"

)

func main() {

	fmt.Println("kuiz satu penjumlahan dengan fungsi diluar main.go")

	hasil := kalkulasi.PanggilPenjumlahan(10, 20)
	fmt.Println("Hasil penjumlahan:", hasil)
}

