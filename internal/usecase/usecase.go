package usecase

import (
	"p1-simulasi-livecode-3-Dedenruslan19/internal/entity"
	"p1-simulasi-livecode-3-Dedenruslan19/internal/repository"
)

type CustomerUsecase interface {
    GenerateReport() ([]entity.FieldReport, error)
    ListCustomersWithoutPayments() ([]entity.CustomerBooking, error) 
}

type customerUsecase struct {
    repo repository.CustomerRepository
}

func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
    return &customerUsecase{repo}
}

func (u *customerUsecase) GenerateReport() ([]entity.FieldReport, error) {
    return u.repo.GetFieldReports()
}

func (u *customerUsecase) ListCustomersWithoutPayments() ([]entity.CustomerBooking, error) {
    return u.repo.GetCustomerBookings()
}