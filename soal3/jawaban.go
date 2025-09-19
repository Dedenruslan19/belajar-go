package main

import (
	"fmt"
	"strconv"
	"strings"
)

// func check untuk memeriksa input ganjil atau genap
func Check() {
	for { // loop untuk inputan sampai valid
		var input string
		fmt.Print("Enter an integer: ")
		fmt.Scanln(&input)

		// coba konversi string ke integer
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue // lanjutkan ke iterasi berikutnya (minta input lagi)
		}

		// jika konversi berhasil, keluar dari loop
		if number%2 == 0 {
			fmt.Printf("%d is even\n", number)
		} else {
			fmt.Printf("%d is odd\n", number)
		}
		break // keluar dari loop
	}
}

func main() {
    // infinite loop
    for {
        // memanggil fungsi check()
        Check()

        // nested loop untuk opsi lanjut atau keluar program
        for {
            // opsi untuk keluar dari program atau tidak
            fmt.Print("Do you want to continue? (yes/no): ")
            var option string
            fmt.Scanln(&option)
            
            // mengonversi input menjadi huruf kecil untuk perbandingan
            lowerCase := strings.ToLower(option)

            // jika inputnya "no", keluar dari program
            if lowerCase == "no" {
                fmt.Println("Goodbye!")
                return
            // jika inputnya "yes", keluar dari loop bersarang lanjutkan loop func check()
            } else if lowerCase == "yes" {
                break
            // jika inputnya tidak valid, cetak pesan kesalahan
            } else {
                fmt.Println("Invalid input, please type 'yes' or 'no'.")
            }
        }
    }
}