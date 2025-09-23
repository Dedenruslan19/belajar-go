package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MenuItem struct {
	Name  string
	Price float64
}

const (
	taxRate = 0.07
	tipRate = 0.20
)

func main() {
	menu := []MenuItem{
		{"Pizza", 10.99},
		{"Salad", 7.99},
		{"Burger", 8.99},
		{"Fries", 2.99},
		{"Soda", 1.99},
		{"Rice", 3.99},
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Hacktiv8 restaurant!")
	fmt.Println("Here's our menu:")
	for _, item := range menu {
		fmt.Printf("%s: $%.2f\n", item.Name, item.Price)
	}

	// order map -> item name : qty
	dataOrder := make(map[string]int)

	for {
		fmt.Print("\nWhat would you like to order? (Type 'done' to finish): ")
		inputFood, _ := reader.ReadString('\n')
		inputFood = strings.TrimSpace(inputFood)

		if strings.ToLower(inputFood) == "done" {
			break
		}

		// cek apakah makanan ada di menu
		var selectedItem *MenuItem
		for _, item := range menu {
			if strings.EqualFold(item.Name, inputFood) {
				selectedItem = &item
				break
			}
		}

		if selectedItem == nil {
			fmt.Println("Sorry, that menu item doesn't exist.")
			continue
		}

		for {
			fmt.Printf("How many %s would you like? ", selectedItem.Name)
			qtyStr, _ := reader.ReadString('\n')
			qtyStr = strings.TrimSpace(qtyStr)
			qty, err := strconv.Atoi(qtyStr)

			if err != nil || qty <= 0 {
				fmt.Println("Invalid quantity, please try again.")
				continue
			}

			dataOrder[selectedItem.Name] += qty
			break
		}
	}

	// hitung total
	fmt.Println("\nYour order summary:")
	var subtotal float64
	for _, item := range menu {
		if qty, ok := dataOrder[item.Name]; ok {
			itemTotal := item.Price * float64(qty)
			fmt.Printf("%s: %d x $%.2f = $%.2f\n", item.Name, qty, item.Price, itemTotal)
			subtotal += itemTotal
		}
	}

	tax := subtotal * taxRate
	tip := subtotal * tipRate
	total := subtotal + tax + tip

	fmt.Printf("\nSubtotal: $%.2f\n", subtotal)
	fmt.Printf("Tax (%.0f%%): $%.2f\n", taxRate*100, tax)
	fmt.Printf("Tip (%.0f%%): $%.2f\n", tipRate*100, tip)
	fmt.Printf("Total: $%.2f\n", total)
}
