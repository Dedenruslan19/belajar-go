package main

import (
	"fmt"
	"strconv"
	"strings"
)

func realInt(n int) {
	if n < 0 {
		fmt.Printf("Angka %d merupakan bilangan negatif\n", n)
	} else if n > 0{
		fmt.Printf("Angka %d merupakan bilangan positif\n", n)
	} else {
		fmt.Printf("Angka anda adalah nol.")
	}
}
func main() {
	for {
		var inputValidation string
		fmt.Print("Cek angka dari:")
		fmt.Scan(&inputValidation)

		num, err := strconv.Atoi(inputValidation)
		if err != nil{
			fmt.Println("Input harus berupa angka.")
			continue
		}
		realInt(num)
		var continueInput string
        fmt.Print("Apakah Anda ingin melanjutkan? (y/n): ")
        fmt.Scan(&continueInput)

        // Periksa pilihan pengguna
        if strings.ToLower(continueInput) != "y" {
            fmt.Println("Program dihentikan.")
            break // Keluar dari loop
        }
	}
}