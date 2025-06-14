package main

import "fmt"

func main() {
	// MULTIPLE RETURN FUNCTION
	keliling, luas := calculateFunctionMultipleReturnValue(10, 5)
	fmt.Println("Keliling:", keliling)
	fmt.Println("Luas:", luas)

	// FUNCTION PREDEFINE RETURN VALUE
	keliling, luas = calculateFunctionPredefinedReturnValue(10, 5)
	fmt.Println("Keliling:", keliling)
	fmt.Println("Luas:", luas)
}

// MULTIPLE RETURN FUNCTION
func calculateFunctionMultipleReturnValue(panjang int, lebar int) (int, int) {
	keliling := 2 * (panjang + lebar)
	luas := panjang * lebar
	return keliling, luas
}

// FUNCTION PREDEFINE RETURN VALUE
func calculateFunctionPredefinedReturnValue(panjang int, lebar int) (keliling int, luas int) {
	keliling = 2 * (panjang + lebar)
	luas = panjang * lebar
	return
}


