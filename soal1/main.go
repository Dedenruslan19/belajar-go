package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// seed random generator dengan time
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// define var untuk sinkronisasi dengan goroutine 
	var wg sync.WaitGroup

	// looping i 1-5
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// menapilkan angka random dengan goroutine
		go func(id int) {
			defer wg.Done()
			num := r.Intn(100)
			fmt.Printf("Goroutine %d menghasilkan angka: %d\n", id, num)
		}(i)
	}

	// tunggu semua goroutine selesai
	wg.Wait()
}
