/*
	interface adalah kontrak yang berisikan method yang akan diimplementasikan kedalam sebuah struct
*/

package main

import (
	"fmt"
)

type KalkulasiGaji interface{
	gaji() int
}

type PegawaiTetap struct {
	GajiPokok int
	Tunjangan int
}

func (p PegawaiTetap) gaji() int {
	return p.GajiPokok + p.Tunjangan 
}

type PegawaiKontrak struct {
	GajiPokok int
}

func (k PegawaiKontrak) gaji() int {
	return k.GajiPokok
}

type Freelance struct{
	RatePerJam int
	JamKerja int
}

func (f Freelance) gaji() int {
	return f.RatePerJam * f.JamKerja
}

func PengeluaranGaji(kalkulasi []KalkulasiGaji) int {
	total := 0
	for _, k := range kalkulasi {
		total += k.gaji()
	}
	return total
}

func main()  {
	pt1 := PegawaiTetap{1,2}
	pt2 := PegawaiTetap{3,4}

	pk1 := PegawaiKontrak{5}
	pk2 := PegawaiKontrak{2}

	f1 := Freelance{2,6}

	totalpengeluarangaji := PengeluaranGaji([]KalkulasiGaji{pk1,pk2,pt1,pt2,f1})

	fmt.Println(totalpengeluarangaji)
}