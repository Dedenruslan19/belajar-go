package main

import (
	"fmt"
	"strconv"
)

// function untuk menentukan bilangan prima
func isPrime(n int) bool {
	// jika bilangan kurang dari sama dengan 1, kembalikan nilai false
	if n <= 1 {
		return false
	}
	// iterasi untuk pengecekan value tersebut bilangan prima atau tidak
	for i := 2; i*i <= n; i++{
		if n%i == 0{
			// jika value tersebut habis dibagi maka return false
			return false
		}
	}
	return true
}

func main() {
	// definisikan var untuk dicek dengan func isPrime()
    var lowerInt int
    var upperInt int

	// infinite loop
    for {
		// definisikan var dari input
        var lowerInput string
        fmt.Print("Enter the lower bound: ")
        fmt.Scan(&lowerInput)

        // handling error jika input berupa string (mengecek string yang diinput apakah ada pada int)
        lower, err1 := strconv.Atoi(lowerInput)
        if err1 != nil {
            fmt.Println("Invalid input. Please enter a valid integer for lower bound.")
            continue
        }
		// jika ada, simpan pada lowerInt untuk dicek
        lowerInt = lower

		// nested loop lanjutan untuk input ke upperInt
        for {
			// definisikan var dari input
            var upperInput string
            fmt.Print("Enter the upper bound: ")
            fmt.Scan(&upperInput)
            
			// handling error jika input berupa string (mengecek string yang diinput apakah ada pada int)
            upper, err2 := strconv.Atoi(upperInput)
            if err2 != nil {
                fmt.Println("Invalid input. Please enter a valid integer for upper bound.")
                continue
            }

            // pengecekan lanjutan upper harus lebih besar dari lower
            if upper < lower {
                fmt.Println("Upper bound must be greater than or equal to lower bound.")
                continue
            }

            // jika value merupakan int, simpan pada upperInt untuk dicek
            upperInt = upper
            break // keluar dari loop upper bound setelah input valid, lanjutkan ke loop lower bound
        }
        break // keluar dari loop lower bound setelah loop upper bound selesai
    }

	// tampilkan ke layar nilai dari lowerInt dan upperInt kemudian panggil fungsi isPrime()
    fmt.Printf("Prime numbers between %d and %d are:\n", lowerInt, upperInt)
    for i := lowerInt; i <= upperInt; i++ {
        if isPrime(i) {
            fmt.Println(i)
        }
    }
}