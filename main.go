package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{
		"Pizza": 10.99,
		"Salad": 7.99,
		"Burger": 8.99,
		"Fries": 2.99,
		"Soda": 1.99,
		"Rice": 3.99,
	}

	fmt.Println("Welcome to Hacktiv8 restaurant")
	fmt.Println("Menu: ")

	for item, value := range menu {
		fmt.Printf("%s - $%.2f\n",item, value)
	}

	var orderTotal float64
	
	// Simpan user input ke map baru
	dataOrder := make(map[string]float64)

	for {
		fmt.Print("What would you like to order? (Type done to finish) \n")

		// Define variable input nama makanan
		var inputFood string
		// Error handling untuk nilai kosong di input nama makanan
		_, err := fmt.Scanln(&inputFood)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		// Break jika input "done" pada inputFood
		if inputFood == "done"{
			break
		}

		// 
		price,exist:=menu[inputFood]
		if !exist {
			fmt.Println("Sorry that menu doesn't exist.")
			continue
		}

		orderTotal += price
		
		fmt.Printf("How many %s would you like?\n", inputFood)
		var inputTotalItem int
		_, err = fmt.Scanln(&inputTotalItem)
		if err != nil {
			fmt.Println("Invalid count item")
			continue
		}
		
		dataOrder[inputFood] = float64(inputTotalItem)

		continue
	}
	fmt.Print("Your order summary:\n")
	for k,v:= range dataOrder{
		fmt.Printf("%s,$%.2f", k, v)
	}
}