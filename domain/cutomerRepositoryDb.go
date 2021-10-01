package domain

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/JyotsnaGorle/banking/app/errs"
	"github.com/JyotsnaGorle/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)

	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)

	}

	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewNotFoundError("Error while querying customer table")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {

	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while scanning customer" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &c, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {

	dbUser := os.Getenv("DB_USER")
	dbAddress := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:@tcp(%s:%s)/%s", dbUser, dbAddress, dbPort, dbName)

	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}

}
