package main

import (
	"bufio"   // untuk input teks dengan spasi
	"fmt"     // untuk print output
	"os"      // untuk akses stdin
	"strconv" // untuk konversi string → int
	"strings" // untuk manipulasi string
)

// Struct untuk menyimpan data buku
type JudulItem struct {
	Judul string   // judul buku
	Price float64  // harga buku
}

// Konstanta untuk pajak dan ongkos kirim
const (
	taxRate  float64 = 0.10 // pajak 10%
	shipRate float64 = 0.15 // ongkos kirim 15%
)

func main() {
	// Daftar buku tersedia
	items := []JudulItem{
		{"Go Programming", 15.50},
		{"Phyton", 12.75},
		{"Data Science 101", 20.00},
		{"Web Development", 18.25},
	}

	// Reader untuk membaca input dari user
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Book Store")
	fmt.Println("--- Book item ---")
	// Tampilkan katalog
	for _, item := range items {
		fmt.Printf("%s: $%.2f\n", item.Judul, item.Price)
	}

	// Map untuk menyimpan pesanan user (judul -> jumlah)
	dataOrder := make(map[string]int)

	// Proses input pesanan
	for {
		fmt.Print("\nWhat would you like to order? (Type 'done' to finish): ")
		inputJudul, _ := reader.ReadString('\n')
		inputJudul = strings.TrimSpace(inputJudul)

		// Jika user mengetik "done", keluar dari loop
		if strings.ToLower(inputJudul) == "done" {
			break
		}

		// Cari buku sesuai input
		var selectedItem *JudulItem
		for _, item := range items {
			if strings.EqualFold(item.Judul, inputJudul) {
				selectedItem = &item
				break
			}
		}

		// Jika buku tidak ditemukan
		if selectedItem == nil {
			fmt.Println("The item doesn't exist.")
			continue
		}

		// Minta jumlah pesanan
		for {
			fmt.Printf("How many %s would you like? ", selectedItem.Judul)
			qtyStr, _ := reader.ReadString('\n')
			qtyStr = strings.TrimSpace(qtyStr)
			qty, err := strconv.Atoi(qtyStr) // konversi string → int

			if err != nil || qty <= 0 {
				fmt.Println("Invalid quantity, please try again.")
				continue
			}
			// Tambahkan ke pesanan
			dataOrder[selectedItem.Judul] += qty
			break
		}
	}

	// Hitung total pesanan
	fmt.Println("\nYour order summary:")
	var subtotal float64
	for _, item := range items {
		if qty, ok := dataOrder[item.Judul]; ok {
			itemTotal := item.Price * float64(qty)
			fmt.Printf("%s: %d x $%.2f = $%.2f\n", item.Judul, qty, item.Price, itemTotal)
			subtotal += itemTotal
		}
	}

	// Hitung pajak dan ongkos kirim
	tax := subtotal * taxRate
	ship := subtotal * shipRate
	total := subtotal + tax + ship

	// Tampilkan ringkasan harga
	fmt.Printf("\nSubtotal: $%.2f\n", subtotal)
	fmt.Printf("Tax (%.0f%%): $%.2f\n", taxRate*100, tax)
	fmt.Printf("Shipping (%.0f%%): $%.2f\n", shipRate*100, ship)
	fmt.Printf("Total: $%.2f\n", total)
}


// PROGRAM BookStore

// DEFINE struct JudulItem:
//     Judul (string)
//     Price (float)

// DEFINE CONST taxRate = 0.10
// DEFINE CONST shipRate = 0.15

// DECLARE list items = [
//     ("Go Programming", 15.50),
//     ("Phyton", 12.75),
//     ("Data Science 101", 20.00),
//     ("Web Development", 18.25)
// ]

// PRINT "Book Store"
// PRINT daftar items

// DECLARE dataOrder = map kosong

// LOOP forever:
//     INPUT inputJudul
//     IF inputJudul == "done" THEN
//         BREAK loop

//     SET selectedItem = NIL
//     FOR setiap item dalam items:
//         IF item.Judul == inputJudul (case-insensitive):
//             selectedItem = item
//             BREAK

//     IF selectedItem == NIL:
//         PRINT "The item doesn't exist"
//         CONTINUE loop

//     LOOP:
//         INPUT qty
//         IF qty <= 0:
//             PRINT "Invalid quantity"
//             CONTINUE loop
//         ELSE
//             TAMBAHKAN qty ke dataOrder[selectedItem.Judul]
//             BREAK loop

// // Hitung total
// subtotal = 0
// FOR setiap item dalam items:
//     IF item ada di dataOrder:
//         itemTotal = item.Price * qty
//         PRINT item, qty, price, itemTotal
//         subtotal += itemTotal

// tax = subtotal * taxRate
// ship = subtotal * shipRate
// total = subtotal + tax + ship

// PRINT subtotal, tax, ship, total
// SELESAI
