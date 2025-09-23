package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Medicine struct {
	Item  string
	Price float64
}

var menu = []Medicine{
	{"Cough Syrup", 49750},
	{"Pain Killers", 129500},
	{"Antibiotics", 225000},
	{"Antacids", 275500},
	{"Vitamins", 285000},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Pharmacy Bill System!")

	var orderTotal float64 = 0

	for {
		// tampilkan menu
		fmt.Println("\nHere are our items:")
		for _, m := range menu {
			fmt.Printf("%s: Rp%.0f\n", m.Item, m.Price)
		}

		// input item
		itemIndex := -1
		for {
			fmt.Print("\nWhat item would you like to order? ")
			barisItem, _ := reader.ReadString('\n')
			barisItem = strings.TrimSpace(barisItem)

			if barisItem == "" {
				fmt.Println("Please input what you need from list.")
				continue
			}

			// cek angka (pilih nomor menu)
			if n, err := strconv.Atoi(barisItem); err == nil {
				if n >= 1 && n <= len(menu) {
					itemIndex = n - 1
					break
				} else {
					fmt.Println("Invalid number. Please choose from the menu.")
					continue
				}
			}

			// cek nama obat
			for i, m := range menu {
				if strings.EqualFold(m.Item, barisItem) {
					itemIndex = i
					break
				}
			}
			if itemIndex == -1 {
				fmt.Println("Invalid input, please input valid medicine name from list.")
				continue
			}
			break
		}

		// input jumlah unit
		units := 0
		maxUnits := 99
		for {
			fmt.Printf("How many units of %s? ", menu[itemIndex].Item)
			rawUnits, _ := reader.ReadString('\n')
			rawUnits = strings.TrimSpace(rawUnits)

			if rawUnits == "" {
				fmt.Println("No input provided. Please enter the number of units.")
				continue
			}

			n, err := strconv.Atoi(rawUnits)
			if err != nil || n <= 0 {
				fmt.Println("Invalid number of units. Please enter a positive number.")
				continue
			}
			if n > maxUnits {
				fmt.Printf("Number of units too large. Please enter less than %d.\n", maxUnits+1)
				continue
			}
			units = n
			break
		}

		// hitung subtotal
		subtotal := menu[itemIndex].Price * float64(units)
		orderTotal += subtotal

		// tanya apakah mau order lagi
		fmt.Print("Would you like to order another item? (yes/no) ")
		again, _:= reader.ReadString('\n')
		again = strings.ToLower(strings.TrimSpace(again))

		if again == "yes" {
			continue
		} else if again == "no" {
			finalPrice := orderTotal

			//diskon
			if orderTotal > 1000000 {
				d := 10
				discount := orderTotal * float64(d) / 100
				finalPrice -= discount
				fmt.Printf("Congratulations you qualify for a %d%% discount. Your total after discount is Rp %.0f\n",d, finalPrice)
			}

			// tax
			tax := finalPrice * 0.08
			finalPrice += tax
			fmt.Printf("With tax applied, your final total is Rp %.0f", tax)

			// get antacids
			if orderTotal > 2000000 {
				fmt.Println("\nAddtionally, you qualify for free unit of Antacids. Enjoy your additional item!")
			}

			return
		} else {
			fmt.Println("Please input a valid option: yes or no.")
		}
	}
}