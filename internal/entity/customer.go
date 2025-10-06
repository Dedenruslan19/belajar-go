package entity

import "time"

type Customer struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
}

type FieldReport struct {
	FieldName     string
	TotalBookings int
	TotalRevenue  float64
}

type CustomerBooking struct {
	CustomerName string
	FieldName string
	BookingID    int
	BookingDate  time.Time
}