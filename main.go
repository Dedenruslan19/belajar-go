package main

import "fmt"

func main() {
	i := 21
	// menampilkan nilai i : 21 fmt.Printf("%T \n", i)contoh : fmt.Printf("%v \n", i)
	fmt.Printf("%v \n", i)

	// menampilkan tipe data dari variabel i
	fmt.Printf("%T \n", i)

	// menampilkan tanda % 
	fmt.Println("%")

	// menampilkan nilai boolean j : true 
	j := true
	fmt.Printf("%t \n", j)

	// menampilkan unicode russia : Я (ya) 
	fmt.Println("\u042F")

	// menampilkan nilai base 10 : 21 
	baseInt := 21
	fmt.Printf("%d \n", baseInt)

	// menampilkan nilai base 8 :25 
	baseInt2 := 25
	fmt.Printf("%d \n", baseInt2)

	// menampilkan nilai base 16 : f
	baseDecimal := 15
	fmt.Printf("%x \n", baseDecimal)

	// menampilkan nilai base 16 : F 13 
	baseDecimal2 := 3859
	fmt.Printf("%X \n", baseDecimal2)

	// menampilkan unicode karakter Я : U+042F
	fmt.Println("Unicode \u042F : U+042F")

	// menampilkan float : 123.456000 
	var k float64 = 123.456
	fmt.Printf ("k : %f\n",k)
	
	// menampilkan float scientific : 1.234560E+02 expected output
	fmt.Printf("Notasi ilmiah : %E\n", k)
}