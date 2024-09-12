package main

import (
	"fmt"
	"math/rand"
)

func main() {
	min := 1
	max := 30
	random := rand.Intn(max-min+1) + min

	var user int

	fmt.Println("tembak angka sampai 1-30")
	fmt.Scan(&user)

	if user < min || user > max {
		fmt.Println("anda salah imput,harap input 1-30")
	} else if user == random {
		fmt.Println("anda benar")
	} else if user < random {
		fmt.Println("nomor terlalu kecil")
	} else {
		fmt.Println("Nomor Terlalu Besar")
	}

}
