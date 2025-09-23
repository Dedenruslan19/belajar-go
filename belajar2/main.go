package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type JudulItem struct {
	Judul string
	Price float64
}

const (
	taxRate  = 0.10
	shipRate = 0.15
)

func main() {
	items := []JudulItem{
		{"Go Programming", 15.50},
		{"Pyton", 12.75},
		{"Sharp", 13.75},
		{"C", 16.50},
		{"Java", 12.70},
	}
	for _, i := range items{
		fmt.Printf("%s: $%.2f\n", i.Judul, i.Price)
	}

	reader := bufio.NewReader(os.Stdin)

	dataOrder := make(map[string]int)

	for {
		fmt.Print("Silahkan pilih judul buku! (atau ketik selesai)")
		inputJudul, _ := reader.ReadString ('\n')
		inputJudul = strings.TrimSpace(inputJudul)

		if strings.ToLower(inputJudul) == "selesai"{
			break
		}
		
		var selectedItem *JudulItem
		for _, i := range items {
			if strings.EqualFold(i.Judul, inputJudul){
				selectedItem = &i
				break
			}
		}

		if selectedItem == nil {
			fmt.Println("Buku tidak ada.")
			continue
		}

		for {
			fmt.Printf("Mau pesan buku %s berapa?", inputJudul)
			inputQty, err := reader.ReadString ('\n')
			inputQty = strings.TrimSpace(inputQty)

			qty, err := strconv.Atoi(inputQty)
			if err != nil {
				fmt.Print("Masukkan angka yang benar.\n")
				continue
			} 
			if qty <= 0 {
				fmt.Print("Invalid quantity.\n")
				continue
			}
			dataOrder[selectedItem.Judul] += qty
			break
		}
	}
	fmt.Println("Orderan anda:")
	var subtotal float64
	for _, i := range items {
		if qty, ok := dataOrder[i.Judul]; ok{
			itemTotal := i.Price * float64(qty)
			fmt.Printf("%s: %d x $%.2f = $%.2f\n", i.Judul, qty, i.Price, itemTotal)
			subtotal += itemTotal
		}
	}

	tax := subtotal * taxRate
	ship := subtotal * shipRate
	total := subtotal + tax + ship

	fmt.Printf("\nSubtotal: $%.2f\n", subtotal)
	fmt.Printf("Tax (%.0f%%): $%.2f\n", taxRate*100, tax)
	fmt.Printf("Shipping (%.0f%%): $%.2f\n", shipRate*100, ship)
	fmt.Printf("Total: $%.2f\n", total)
}