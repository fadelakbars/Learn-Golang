package main

import "fmt"

func main()  {
	title := "golang the best programming language"

	for index, huruf := range title {
		switch huruf {
		case 'a', 'i', 'u', 'e', 'o':
			if index % 2 == 0 {
				fmt.Printf("Huruf vokal '%c' pada indeks genap %d\n", huruf, index)
			} 
		default:
		}
	}
}