package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"p1-simulasi-livecode-3-Dedenruslan19/config"
	"p1-simulasi-livecode-3-Dedenruslan19/internal/delivery/cli"
	"p1-simulasi-livecode-3-Dedenruslan19/internal/repository"
	"p1-simulasi-livecode-3-Dedenruslan19/internal/usecase"
	"strconv"
	"strings"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	fmt.Print("==============================================\n")
	fmt.Print("RONALDO'S MINI SOCCER FIELD BOOKING ADMIN CLI\n")
	fmt.Print("==============================================\n")

	repo := repository.NewCustomerRepository(db)
	uc := usecase.NewCustomerUsecase(repo)

	reader := bufio.NewReader(os.Stdin)

	for {
		// Tampilkan Menu
		fmt.Print("\n===== MENU =====\n")
		fmt.Println("1. Generate Revenue Report")
		fmt.Println("2. List Customers Without Payments")
		fmt.Println("3. Exit")
		fmt.Print("Pilih Opsi (1-3): ")

		// 🆕 Baca input sebagai string hingga karakter newline
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}

		// Bersihkan dan konversi input
        // TrimSpace untuk menghapus spasi/newline di awal/akhir input.
		trimmedInput := strings.TrimSpace(input)
		choice, err := strconv.Atoi(trimmedInput) // Konversi string ke integer

		if err != nil {
			fmt.Println("Pilihan tidak valid. Silakan masukkan angka.")
			continue
		}

		switch choice {
		case 1:
			cli.ShowFieldReports(uc)

		case 2:
			cli.ShowCustomerBookings(uc)

		case 3:
			fmt.Println("Exiting program. Goodbye!")
			os.Exit(0)

		default:
			fmt.Println("Pilihan tidak valid. Silakan masukkan 1, 2, atau 3.")
		}
	}
}