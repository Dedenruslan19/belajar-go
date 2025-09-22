package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{
		"Pizza":  10.99,
		"Salad":  7.99,
		"Burger": 8.99,
		"Fries":  2.99,
		"Soda":   1.99,
		"Rice":   3.99,
	}

	fmt.Println("Welcome to Hacktiv8 restaurant!")
	fmt.Println("Here's our menu:")
	for item, value := range menu {
		fmt.Printf("%s: $%.2f\n", item, value)
	}

	// Simpan user input ke map baru (item -> jumlah)
	dataOrder := make(map[string]int)

	for {
		fmt.Print("What would you like to order? (Type 'done' to finish)\n")

		var inputFood string
		_, err := fmt.Scanln(&inputFood)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if inputFood == "done" {
			break
		}

		// cek apakah makanan ada di menu
		_, exist := menu[inputFood]
		if !exist {
			fmt.Println("Sorry, that menu item doesn't exist.")
			continue
		}

		for {
			fmt.Printf("How many %s would you like?\n", inputFood)
			var qty int
			_, err = fmt.Scanln(&qty)
			if err != nil || qty <= 0 {
				fmt.Println("Invalid quantity")
				continue
			}
			dataOrder[inputFood] += qty
			break
		}
	}

	// hitung total
	fmt.Println("\nYour order summary:")
	var subtotal float64
	for item, qty := range dataOrder {
		price := menu[item]
		itemTotal := price * float64(qty)
		fmt.Printf("%s: %d x $%.2f = $%.2f\n", item, qty, price, itemTotal)
		subtotal += itemTotal
	}

	tax := subtotal * 0.07
	tip := subtotal * 0.20
	total := subtotal + tax + tip

	fmt.Printf("Subtotal: $%.2f\n", subtotal)
	fmt.Printf("Tax (7%%): $%.2f\n", tax)
	fmt.Printf("Tip (20%%): $%.2f\n", tip)
	fmt.Printf("Total: $%.2f\n", total)
}
