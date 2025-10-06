package repository

import (
	"database/sql"
	"p1-simulasi-livecode-3-Dedenruslan19/internal/entity"
)

type CustomerRepository interface {
	GetAll() ([]entity.Customer, error)
	GetFieldReports() ([]entity.FieldReport, error)
	GetCustomerBookings() ([]entity.CustomerBooking, error)
}

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) GetAll() ([]entity.Customer, error) {
	rows, err := r.db.Query(`SELECT customerid, firstname, lastname, email, phonenumber FROM customers`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []entity.Customer
	for rows.Next() {
		var c entity.Customer
		if err := rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email, &c.PhoneNumber); err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

// 🆕 Function baru: Menampilkan fieldname, total bookings, dan total revenue
func (r *customerRepository) GetFieldReports() ([]entity.FieldReport, error) {
	query := `
		SELECT 
			f.fieldname,
			COUNT(b.bookingid) AS total_bookings,
			(COUNT(b.bookingid) * f.hourlyrate) AS total_revenue
		FROM fields f
		LEFT JOIN bookings b ON f.fieldid = b.fieldid
		GROUP BY f.fieldid, f.fieldname, f.hourlyrate
		ORDER BY total_revenue DESC;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []entity.FieldReport
	for rows.Next() {
		var rpt entity.FieldReport
		if err := rows.Scan(&rpt.FieldName, &rpt.TotalBookings, &rpt.TotalRevenue); err != nil {
			return nil, err
		}
		reports = append(reports, rpt)
	}

	return reports, nil
}

func (r *customerRepository) GetCustomerBookings() ([]entity.CustomerBooking, error) {
	query := `
		SELECT 
			CONCAT(c.firstname, ' ', c.lastname) AS customername,
			f.fieldname,
			b.bookingid,
			b.bookingdate
		FROM bookings b
		JOIN customers c ON b.customerid = c.customerid
		JOIN fields f ON b.fieldid = f.fieldid
		ORDER BY b.bookingdate DESC;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []entity.CustomerBooking
	for rows.Next() {
		var rpt entity.CustomerBooking
		if err := rows.Scan(&rpt.CustomerName, &rpt.FieldName, &rpt.BookingID, &rpt.BookingDate); err != nil {
			return nil, err
		}
		bookings = append(bookings, rpt)
	}

	return bookings, nil
}




