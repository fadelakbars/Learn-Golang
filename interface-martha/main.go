/*
	interface digunakna agar kode bisa menerima banyak jenis struct (polimorphism)
*/

package main

import "fmt"

type HitungPersegi interface {
	luas() int
	keliling() int
}

type Persegi struct {
	Sisi int
}

func (s Persegi) luas() int {
	return s.Sisi * s.Sisi
}

func (s Persegi) keliling() int{
	return s.Sisi * 4
}

func main()  {
	var p HitungPersegi
	p = Persegi{6}
	keliling := p.keliling()
	luas := p.luas()

	fmt.Println(keliling)
	fmt.Println(luas)

	q := Persegi{3}
	kelilingq := q.keliling()
	luasq := q.luas()

	fmt.Println(kelilingq)
	fmt.Println(luasq)
}