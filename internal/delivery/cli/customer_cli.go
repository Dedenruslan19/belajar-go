package cli

import (
	"fmt"
	"p1-simulasi-livecode-3-Dedenruslan19/internal/entity"
	"p1-simulasi-livecode-3-Dedenruslan19/internal/usecase"
)

func PrintCustomerBookings(bookings []entity.CustomerBooking) {
	if len(bookings) == 0 {
		fmt.Println("\nTidak ada data booking pelanggan.")
		return
	}

	fmt.Println("\n====================================================================")
	fmt.Println("                    DAFTAR BOOKING PELANGGAN")
	fmt.Println("====================================================================")
	fmt.Printf("%-5s | %-20s | %-12s\n", "ID", "CUSTOMER", "TANGGAL")
	fmt.Println("--------------------------------------------------------------------")

	for _, b := range bookings {
		fmt.Printf("%-5d | %-20s | %-15s\n",
			b.BookingID,
			b.CustomerName,
			b.BookingDate.Format("2006-01-02"),
		)
	}

	fmt.Println("--------------------------------------------------------------------")
	fmt.Printf("Total bookings: %d\n", len(bookings))
	fmt.Println("====================================================================")
}

func ShowFieldReports(uc usecase.CustomerUsecase) {
    // ➡️ Memanggil Use Case
    reports, err := uc.GenerateReport() 
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("=== Field Reports ===")
    for _, r := range reports {
        fmt.Printf("%-15s | %3d bookings | Rp %.2f\n", r.FieldName, r.TotalBookings, r.TotalRevenue)
    }
}

// 🆕 Menerima usecase.CustomerUsecase
func ShowCustomerBookings(uc usecase.CustomerUsecase) {
    // ➡️ Memanggil Use Case
    bookings, err := uc.ListCustomersWithoutPayments() 
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    PrintCustomerBookings(bookings)
}