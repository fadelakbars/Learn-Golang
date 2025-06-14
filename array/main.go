package main

import "fmt"

func main()  {
	arr := []int{1, 2, 3, 4, 5}

	// MENGAKSES ELEMEN ARRAY
	fmt.Println("\n",arr[1])

	// MENGUBAH ELEMEN ARRAY
	arr[1] = 10
	fmt.Println("Setelah mengubah elemen kedua:", arr[1  ])

	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}