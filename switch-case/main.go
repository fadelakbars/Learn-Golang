package main

import "fmt"

func main()  {
	
	// REGULER SWITCH CASE
	hari := "mei"
	switch hari {
	case "jumat":
		fmt.Println("ibadah")
	case "sabtu", "minggu":
		fmt.Println("libur")
	default:
		fmt.Println("hari kerja")
	}
}