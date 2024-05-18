package db

import (
	"database/sql"
	"ekart/model"
	"fmt"
	"log"
	"time"
)

func CreateCustomerTables(*sql.DB) {
	db := DbInIt()

	query := `CREATE TABLE IF NOT EXISTS  customers (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		email VARCHAR(255),
		mobile VARCHAR(20),
		password VARCHAR(255),
		gender VARCHAR(10),
		adult BOOLEAN,
		created_at TIMESTAMP
	)`
	res, err := db.Exec(query)
	if err != nil {
		log.Fatalf("error while excute create table query %v", err)
	}
	if res != nil {
		println("Table create succesfuly")
	}

}
func CreateAddressTable(*sql.DB) {
	db := DbInIt()

	query := `CREATE TABLE  IF NOT EXISTS addresses(
		id SERIAL PRIMARY KEY,
		cust_id INT,
		house_no VARCHAR(50) NOT NULL,
		town VARCHAR(100) NOT NULL,
		district VARCHAR(100) NOT NULL,
		state VARCHAR(100) NOT NULL,
		pincode INT,
		FOREIGN KEY (cust_id) REFERENCES customers(id)
	)`
	res, err := db.Exec(query)
	if err != nil {
		log.Fatalf("error while excute create address table query %v", err)
	}
	if res != nil {
		println("Table create succesfuly")
	}
}

func NewCustomer(customer model.Customer) int {
	db := DbInIt()
	defer db.Close()
	CreateCustomerTables(db)
	// Insert customer
	var id int
	err := db.QueryRow(`INSERT INTO customers (name, email, mobile, password, gender, adult, created_at)
	VALUES ($1, $2, $3, $4, $5, $6,$7)
	RETURNING id`,
		customer.Name, customer.Email, customer.Mobile, customer.Password, customer.Gender, customer.Adult, time.Now()).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	// Insert address

	return id

}
func InsertAdd(customer model.Customer) {
	db := DbInIt()
	CreateAddressTable(db)
	var id int
	res, err := db.Exec(`INSERT INTO addresses (cust_id, house_no, town, district, state, pincode)
	VALUES ($1, $2, $3, $4, $5, $6)`,
		id, customer.Address.HouseNo, customer.Address.Town, customer.Address.District, customer.Address.State, customer.Address.Pincode)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		fmt.Println(res)
	}
}
