package domain

import (
	"github.com/Altamashattari/banking-application/dto"
	"github.com/Altamashattari/banking-application/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	// status == 1 - Active status == 0 - Inactive status == "" - No filter
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

func (c Customer) statusAsText() (statusAsText string) {
	statusAsText = "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}
