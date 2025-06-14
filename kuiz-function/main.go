package main

import (
	"fmt"
)

func main()  {

	angka := []int{1, 2, 3, 4, 5}
	total := sum(angka)
	fmt.Println("Total hasil penjumlahan :", angka ," adalah ", total)
	
	calculateResult, err := calculate(10, 0, "/")
	if calculateResult >= 0 && err == nil {
		fmt.Println("Hasil penjumlahan = ", calculateResult)
	}else if err != nil {
		fmt.Println("Error:", err)
	}
}

func sum(number []int) (total int)  {
	total = 0
	for _, angka := range number {
		total += angka
	}
	return
}

func calculate(angka1 int, angka2 int, operator string) (int, error) {
	var err error
	hasil := 0

	switch operator {
	case "+":
		hasil = angka1 + angka2
	case "-":
		hasil = angka1 - angka2
	case "*":
		hasil = angka1 * angka2
	case "/":
		if angka2 != 0 {
			hasil = angka1 / angka2
		} else {
			err = fmt.Errorf("bilangan tidak boleh dibagi nol")
		}
	default:
		err = fmt.Errorf("operator tidak dikenali")
	}

	return hasil, err
}
