package db

import (
	"database/sql"
	"ekart/model"
	"log"
)

func CreateCustomerTables(*sql.DB) {
	db := DbInIt()
	defer db.Close()
	query := `CREATE TABLE  IF NOT EXISTS customers(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	mobile TEXT NOT NULL,
	password TEXT NOT NULL,
	gender TEXT NOT NULL,
	adult BOOLEAN NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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
	defer db.Close()
	query := `CREATE TABLE  IF NOT EXISTS addresses(
		cust_id INTEGER,
		town TEXT NOT NULL,
		district TEXT NOT NULL,
		state TEXT NOT NULL,
		pincode INTEGER NOT NULL
	)`
	res, err := db.Exec(query)
	if err != nil {
		log.Fatalf("error while excute create address table query %v", err)
	}
	if res != nil {
		println("Table create succesfuly")
	}
}

func NewCustomer(customer model.Customer, address model.Address) int {
	db := DbInIt()

	CreateCustomerTables(db)
	var id int
	query := (`INSERT INTO customers(name,email,mobile,password,gender,adult)VALUES($1,$2,$3,$4,$5,$6)RETURNING id`)
	err := db.QueryRow(query, customer.Name, customer.Email, customer.Mobile, customer.Password, customer.Gender, customer.Adult).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	CreateAddressTable(db)
	db.Query("INSERT INTO addresses (cust_id, house_no, town, district, state, pincode)VALUES($1, $2, $3, $4, $5, $6)", id, address.HouseNo, address.Town, address.District, address.State, address.Pincode)
	return id

}

// func SaveCustomer(customer model.Customer, address model.Address) error {
// 	db := DbInIt()
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	defer func() {
// 		if p := recover(); p != nil {
// 			tx.Rollback()
// 			err = fmt.Errorf("panic: %v", p)
// 		} else if err != nil {
// 			tx.Rollback()
// 		} else {
// 			err = tx.Commit()
// 		}
// 	}()

// 	// Insert customer data
// 	CreateCustomerTables(db)
// 	var id int
// 	query := (`INSERT INTO customers(name,email,mobile,password,gender,adult)VALUES($1,$2,$3,$4,$5,$6) `)
// 	db.QueryRow(query, customer.Name, customer.Email, customer.Mobile, customer.Password, customer.Gender, customer.Adult)

// 	CreateAddressTable(db)
// 	db.Query("INSERT INTO addresses (cust_id, house_no, town, district, state, pincode)VALUES($1, $2, $3, $4, $5, $6)", id, address.HouseNo, address.Town, address.District, address.State, address.Pincode)
// 	// return id

// 	return nil
// }
