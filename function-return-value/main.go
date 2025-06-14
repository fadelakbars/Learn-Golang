package main

import "fmt"

func main()  {
	fmt.Println(kembalkaString("fungsi 1"))
	cetaklangsung("fungsi 2")
}

// FUNGSI ME RETURN STRING
func kembalkaString(jenis string) string {
	return jenis + " ini fmt-nya di main"
}

// FUNGSI YANG MENCETEKA
func cetaklangsung(anu string) {
	fmt.Println(anu, "cukup panggil nama fungsinya di main")
}
