package main

import "fmt"

// function untuk menentukan bilangan prima
func isPrime(n int) bool {
	// jika bilangan di bawah 2, kembalikan nilai false
	if n < 2 {
		return false
	}
	// iterasi untuk pengecekan value tersebut bilangan prima atau tidak 
	for i := 2; i*i <= n; i++{
		// jika value tersebut habis dibagi maka return false
		if n % i == 0{
			return false
		}
	}
	return true
}

func main() {
	// definisikan variable grade sebagai float64
	var lowerInt int
	var upperInt int
	// print untuk input user
	fmt.Print("Enter the lower bound: ")
	// definisikan variable err1 untuk handling error pada lowerInt
	// gunakan pointer untuk merubah value dari lowerInt
	_, err1 := fmt.Scanln(&lowerInt)

	// validasi input lowerInt, 
	// jika input berupa string / input bukan bilangan prima maka tampilkan Invalid input, 
	// jika valid lanjutkan proses
	if err1 != nil || lowerInt < 0 {
		fmt.Println("Invalid input")
		return
	}

	// definisikan variable err1 untuk handling error pada upperInt
	// gunakan pointer untuk merubah value dari upperInt
	fmt.Print("Enter the upper bound: ")
	_, err2 := fmt.Scanln(&upperInt)

	// validasi input upperInt, 
	// jika input berupa string / input tidak lebih besar dari lowerInt maka tampilkan invalid input, 
	// jika valid lanjutkan proses
	if err2 != nil || lowerInt > upperInt {
		fmt.Println("Invalid input")
	}

	
	fmt.Printf("Prime numbers between %d and %d are: \n", lowerInt, upperInt)
	// perulangan untuk menampilkan setiap value yang merupakan bilangan prima
	for i := lowerInt; i <= upperInt ; i++{
		if isPrime(i){
			fmt.Print(i, "\n")
		}
	}
}
