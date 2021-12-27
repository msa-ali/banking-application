package domain

import (
	"database/sql"

	"github.com/Altamashattari/banking-application/errs"
	"github.com/Altamashattari/banking-application/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	// var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSQL)
		// rows, err = d.client.Query(findAllSQL)
	} else {
		findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSQL, status)
		// rows, err = d.client.Query(findAllSQL, status)
	}

	if err != nil {
		logger.Error("Error querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB Error")
	}

	// err = sqlx.StructScan(rows, &customers)

	// if err != nil {
	// 	logger.Error("Error querying customer table " + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected DB Error")
	// }

	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	// 	if err != nil {
	// 		logger.Error("Error while scanning customers " + err.Error())
	// 		return nil, errs.NewUnexpectedError("Unexpected DB Error")
	// 	}
	// 	customers = append(customers, c)
	// }
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var c Customer
	err := d.client.Get(&c, customerSQL, id)

	// row := d.client.QueryRow(customerSQL, id)
	// err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
