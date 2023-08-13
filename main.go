package main

import (
	"bufio"
	"database/sql"
	"enigma-laundry/entity"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "enigmalaundry"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected!")
	}
	return db
}

func main() {
	//menu console
	for {
		fmt.Println("=== Enigma Laundry ===")
		fmt.Println("1. Tambah Pelanggan")
		fmt.Println("2. Ubah Data Pelanggan")
		fmt.Println("3. Hapus Data Pelanggan")
		fmt.Println("4. Tambah Pesanan")
		fmt.Println("5. Ubah Data Pesanan")
		fmt.Println("6. Hapus Data Pesanan")
		fmt.Println("7. Cari Pelanggan")
		fmt.Println("8. Cari Pesanan dan Status")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih tindakan (1/2/3/4/5/6/7/8/9): ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			addCustomer(entity.Customer{})
		case "2":
			updateCustomer(entity.Customer{})
		case "3":
			deleteCustomer(entity.Customer{})
		// case "4":
		// 	addOrder()
		// case "5":
		// 	updateOrder()
		// case "6":
		// 	deleteOder()
		// case "7":
		// 	searchCustomerBy()
		// case "8":
		// 	searchOrderBy()
		case "9":
			fmt.Println("Anda keluar dari aplikasi !")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// add customer
func addCustomer(customer entity.Customer) {
	db := connectDb()
	defer db.Close()
	var err error

	fmt.Println("=== PENDAFTARAN CUSTOMER ===")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Id : ")
	customer.Id, _ = reader.ReadString('\n')
	customer.Id = strings.TrimSpace(customer.Id)

	fmt.Print("Masukkan Nama : ")
	customer.Name, _ = reader.ReadString('\n')
	customer.Name = strings.TrimSpace(customer.Name)

	fmt.Print("Masukkan Alamat : ")
	customer.Address, _ = reader.ReadString('\n')
	customer.Address = strings.TrimSpace(customer.Address)

	fmt.Print("Masukkan Nomor Telepon: ")
	customer.PhoneNumber, _ = reader.ReadString('\n')
	customer.PhoneNumber = strings.TrimSpace(customer.PhoneNumber)

	fmt.Print("Masukkan Email: ")
	customer.Email, _ = reader.ReadString('\n')
	customer.Email = strings.TrimSpace(customer.Email)

	sqlStatement := "INSERT INTO customers (cust_id, cust_name, phone_number, address, email) VALUES ($1, $2, $3, $4, $5);"
	_, err = db.Exec(sqlStatement, customer.Id, customer.Name, customer.PhoneNumber, customer.Address, customer.Email)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Insert Data!")
	}
}

// update customer
func updateCustomer(customer entity.Customer) {
	db := connectDb()
	defer db.Close()
	var err error

	fmt.Println("=== UBAH DATA CUSTOMER ===")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Id : ")
	customer.Id, _ = reader.ReadString('\n')
	customer.Id = strings.TrimSpace(customer.Id)

	fmt.Print("Masukkan Nama : ")
	customer.Name, _ = reader.ReadString('\n')
	customer.Name = strings.TrimSpace(customer.Name)

	fmt.Print("Masukkan Alamat : ")
	customer.Address, _ = reader.ReadString('\n')
	customer.Address = strings.TrimSpace(customer.Address)

	fmt.Print("Masukkan Nomor Telepon: ")
	customer.PhoneNumber, _ = reader.ReadString('\n')
	customer.PhoneNumber = strings.TrimSpace(customer.PhoneNumber)

	fmt.Print("Masukkan Email: ")
	customer.Email, _ = reader.ReadString('\n')
	customer.Email = strings.TrimSpace(customer.Email)

	sqlStatement := "UPDATE customers SET cust_name = $2, phone_number = $3, address = $4, email = $5 WHERE cust_id = $1;"

	_, err = db.Exec(sqlStatement, customer.Id, customer.Name, customer.PhoneNumber, customer.Address, customer.Email)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Update Data!")
	}
}

// delete customer
func deleteCustomer(customer entity.Customer) {
	db := connectDb()
	defer db.Close()
	var err error

	fmt.Println("=== HAPUS DATA CUSTOMER ===")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Id : ")
	customer.Id, _ = reader.ReadString('\n')
	customer.Id = strings.TrimSpace(customer.Id)

	sqlStatement := "DELETE FROM customers WHERE cust_id = $1;"

	_, err = db.Exec(sqlStatement, customer.Id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Delete Data!")
	}
}
