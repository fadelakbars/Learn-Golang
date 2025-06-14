package main

import (
	"fmt"
)

func main()  {
	// REGULER IF ELSE
	var umur int = 18
	if umur >= 17 {
		fmt.Println("boleh membuat sim")
	} else {
		fmt.Println("tidak boleh bawa motor")
	}

	// SHORT STATEMENT IF
	if nilai := 70; nilai >= 75 {
		fmt.Println("gacor kang")
	} else {
		fmt.Println("anda kurang beruntung")
	}

	// PENGELOMPOKAN KONDISI
	status := "Mahasiswa aktif"
	semester := 5
	if (status == "Mahasiswa aktif" && semester >= 6) || status == "Fresh Graduate" {
		fmt.Println("Memenuhi syarat magang")
	} else {
		fmt.Println("Tidak memenuhi syarat magang")
	}


}