package main

import "fmt"

/*
	by default, slice dan map sudah bersifat seperti pointer
*/
func ubahdataslice(data []int)  {
	data[0] = 100
}

func main()  {
	data := []int{1,2,3,4}
	fmt.Println(data)
	ubahdataslice(data)
	fmt.Println(data)
}