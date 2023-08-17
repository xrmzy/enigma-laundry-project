package main

import (
	"bufio"
	"database/sql"
	"enigma-laundry/entity"
	"fmt"
	"math/rand"
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
		case "7":
		searchMenu:
			for {
				fmt.Println("1. Cari berdasarkan ID")
				fmt.Println("2. Cari semua data")
				fmt.Println("0. Kembali ke menu utama")
				fmt.Print("Pilih tindakan (1/2/0) :")

				var searchChoice string
				fmt.Scanln(&searchChoice)
				switch searchChoice {
				case "1":
					searchBy(entity.Customer{})
				case "2":
					searchAll(entity.Customer{})
				case "0":
					break searchMenu // balik ke menu utama
				default:
					fmt.Println("Pilihan tidak valid.")
				}
			}
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

func generateRandomNumber(max int) int {
	return rand.Intn(max)
}

// func generateCustomerId() string {
// 	codePrefix := "CS"
// 	randomNumber := generateRandomNumber(1000)
// 	return fmt.Sprintf("%s%03d", codePrefix, randomNumber)
// }

// func getInput(prompt string) string {
// 	reader := bufio.Newreader(os.Stdin)
// 	fmt.Print(prompt)
// 	inpit, _ := reader.ReadString('\n')
// 	return strings.TrimSpace(input)
// }

// add customer
func addCustomer(customer entity.Customer) {
	db := connectDb()
	defer db.Close()
	var err error

	fmt.Println("=== PENDAFTARAN CUSTOMER ===")
	reader := bufio.NewReader(os.Stdin)

	randomNumber := generateRandomNumber(1000)
	codePrefix := "CS"

	customer.Id = fmt.Sprintf("%s%03d", codePrefix, randomNumber)

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

	//validasi email
	var emailExists bool
	row := db.QueryRow("SELECT EXISTS (SELECT 1 FROM customers WHERE email = $1)", customer.Email)
	err = row.Scan((&emailExists))
	if err != nil {
		panic(err)
	}

	if emailExists {
		fmt.Println("Email sudah digunakan. Silahkan gunakan email lain!")
		return
	}

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

// func addOrder() {

// }

// search customer by ID
func searchBy(customer entity.Customer) {
	db := connectDb()
	defer db.Close()

	fmt.Println("=== CARI PELANGGAN BERDASARKAN ID ===")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Pelanggan: ")
	custID, _ := reader.ReadString('\n')
	custID = strings.TrimSpace(custID)

	sqlStatement := "SELECT cust_id, cust_name, phone_number, address, email FROM customers WHERE cust_id = $1;"
	row := db.QueryRow(sqlStatement, custID)

	err := row.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address, &customer.Email)
	if err == sql.ErrNoRows {
		fmt.Println("Pelanggan tidak ditemukan.")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("ID: %s, Nama: %s, Telepon: %s, Alamat: %s, Email: %s\n",
			customer.Id, customer.Name, customer.PhoneNumber, customer.Address, customer.Email)
	}
}

// search all customers
func searchAll(customer entity.Customer) {
	db := connectDb()
	defer db.Close()

	fmt.Println("=== CARI SEMUA PELANGGAN ===")

	rows, err := db.Query("SELECT cust_id, cust_name, phone_number, address, email FROM customers;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address, &customer.Email)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		fmt.Printf("ID: %s, Nama: %s, Telepon: %s, Alamat: %s, Email: %s\n",
			customer.Id, customer.Name, customer.PhoneNumber, customer.Address, customer.Email)
	}
}
