package main

import (
	"fmt"
)

func main() {
	// definisikan variable grade sebagai float64
	var grade float64
	// print untuk input user
	fmt.Print("Masukkan nilai (0-100): ")
	// definisikan variable err untuk handling error
	// gunakan pointer untuk merubah value dari grade
	_, err := fmt.Scanln(&grade)

	// validasi input dan handling error 
	// jika grade 1-100 maka lanjutkan proses
	if err != nil || grade < 0 || grade > 100 {
		fmt.Println("Invalid grade input")
		return
	}

	// menentukan case berdasarkan value grade dengan default F:Fail
	switch {
	case grade >= 90:
		fmt.Println("A: Excellent")
	case grade >= 80:
		fmt.Println("B: Good")
	case grade >= 70:
		fmt.Println("C: Satisfactory")
	case grade >= 60:
		fmt.Println("D: Needs Improvement")
	default:
		fmt.Println("F: Fail")
	}
}
