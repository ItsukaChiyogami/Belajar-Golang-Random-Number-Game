package main

import (
	"fmt"
	"math/rand"
)

func main() {

	heal := 10

	min := 1
	max := 30
	random := rand.Intn(max-min+1) + min

	var user int

	fmt.Println("tembak angka sampai 1-30")
	for heal > 0 {
		fmt.Printf("Anda memiliki %d heal tersisa\n", heal)
		fmt.Print("Masukkan tebakan: ")
		fmt.Scan(&user)

		if user < min || user > max {
			heal--
			fmt.Println("anda salah imput,harap input 1-30")
		} else if user == random {
			fmt.Println("anda benar")
		} else if user < random {
			heal--
			fmt.Println("nomor terlalu kecil")
		} else {
			heal--
			fmt.Println("Nomor Terlalu Besar")
		}
		if heal == 0 {
			fmt.Println("game over")
		}
	}
}
