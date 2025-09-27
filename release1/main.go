package main

import (
	"fmt"  
	"sync" 
)

// runTask menjalankan serangkaian tugas secara berurutan
func runTask(taskName string, n int) {
	var wg sync.WaitGroup 
	// Membuat channel ber-buffer 1 sebagai 'token' giliran
	ch := make(chan int, 1) 
	// Memulai buffer dari 1
	ch <- 1

	// Perulangan untuk membuat 'n' goroutine/pekerja
	for i := 1; i <= n; i++ {
		// Menambah goroutines setiap ada perulangan
		wg.Add(1)
		// Luncurkan goroutine baru (pekerja)
		go func(i int) {
			defer wg.Done()
			// Perulangan tak terbatas untuk mencoba mendapatkan giliran
			for {
				val := <-ch // Ambil token giliran dari channel (blokir sampai ada)
				if val == i { // Cek apakah token (val) cocok dengan nomor urut goroutine (i)
					fmt.Printf("%s: %d\n", taskName, i) // Cetak output (eksekusi tugas)
					ch <- i + 1 // Kirim token giliran berikutnya (i + 1)
					break // Keluar dari perulangan for (goroutine selesai)
				} else {
					ch <- val // Jika giliran bukan miliknya, kembalikan token ke channel
				}
			}
		}(i) // Panggil/jalankan fungsi anonim dengan nilai i saat ini
	}
	wg.Wait() // Tunggu sampai semua goroutine (worker) memanggil wg.Done()
	close(ch) // Tutup channel setelah semua tugas selesai
}

func main() {
	runTask("Task A", 5) 
	runTask("Task B", 5) 
	runTask("Task C", 5) 
}