package main

import (
	"bufio"
	"database/sql"
	"enigma-laundry/entity"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

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
		case "4":
			addOrder(entity.Orders{})
		case "5":
			updateOrder(entity.Orders{})
		case "6":
			deleteOrder(entity.Orders{})
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
		case "8":
		searchMenuTx:
			for {
				fmt.Println("1. Cari berdasarkan ID Transaksi")
				fmt.Println("2. Cari semua data")
				fmt.Println("0. Kembali ke menu utama")
				fmt.Print("Pilih tindakan (1/2/0) :")

				var searchChoice string
				fmt.Scanln(&searchChoice)
				switch searchChoice {
				case "1":
					searchOrderBy(entity.Orders{})
				case "2":
					SearchAllOrders(entity.Orders{})
				case "0":
					break searchMenuTx // balik ke menu utama
				default:
					fmt.Println("Pilihan tidak valid.")
				}
			}
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

	if isEmailExists(db, customer.Email) {
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

	fmt.Print("Masukkan ID Customer : ")
	customer.Id, _ = reader.ReadString('\n')
	customer.Id = strings.TrimSpace(customer.Id)

	if !isCustomerExists(db, customer.Id) {
		fmt.Println("Pelanggan tidak ditemukan!")
		return
	}

	// row := db.QueryRow("SELECT cust_name, address, phone_number, email FROM customers WHERE cust_id = $1;", customer.Id)
	// err = row.Scan(&customer.Name, &customer.Address, &customer.PhoneNumber, &customer.Email)
	// if err == sql.ErrNoRows {
	// 	fmt.Println("Pelanggan tidak ditemukan!")
	// 	return
	// } else if err != nil {
	// 	panic(err)
	// }

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
	fmt.Print("Masukkan ID Customer : ")
	customer.Id, _ = reader.ReadString('\n')
	customer.Id = strings.TrimSpace(customer.Id)

	if !isCustomerExists(db, customer.Id) {
		fmt.Println("Pelanggan tidak ditemukan!")
		return
	}

	sqlStatement := "DELETE FROM customers WHERE cust_id = $1;"

	_, err = db.Exec(sqlStatement, customer.Id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Delete Data!")
	}
}

// global func untuk getCustomer ID dari Nama
func getCustomerIdByName(db *sql.DB, customerName string) (string, error) {
	var customerId string
	row := db.QueryRow("SELECT cust_id FROM customers WHERE cust_name = $1", customerName)
	err := row.Scan(&customerId)
	if err != nil {
		return "", err
	}
	return customerId, nil
}

// global func untuk getCustomer Name dari ID
func getCustomerNameById(db *sql.DB, customerId string) (string, error) {
	var customerName string
	row := db.QueryRow("SELECT cust_name FROM customers WHERE cust_id = $1", customerId)
	err := row.Scan(&customerName)
	if err != nil {
		return "", err
	}
	return customerName, nil
}

// add order
func addOrder(order entity.Orders) {
	db := connectDb()
	defer db.Close()
	var err error

	fmt.Println("=== TAMBAH TRANSAKSI ===")
	reader := bufio.NewReader(os.Stdin)

	randomNumber := generateRandomNumber(1000)
	codePrefix := "TX"

	order.OrderId = fmt.Sprintf("%s%03d", codePrefix, randomNumber)

	fmt.Print("Masukkan Nama :")
	order.CustomerName, _ = reader.ReadString('\n')
	order.CustomerName = strings.TrimSpace(order.CustomerName)

	// Validasi apakah pelanngan dengan nama tersebut ada
	customerId, err := getCustomerIdByName(db, order.CustomerName)
	if err != nil {
		fmt.Println("Data Pelanggan tidak ditemukan, Silahkan Tambah Pelanggan Terlebih Dahulu!")
		return
	}
	order.CustomerId = customerId
	customerName, err := getCustomerNameById(db, customerId)
	if err != nil {
		panic(err)
	}
	order.CustomerName = customerName

	// Pilih Paket
	fmt.Println("Pilih Paket:")
	fmt.Println("1. PAKET A/BERSIH AMAN")
	fmt.Println("2. PAKET B/BERSIH TENANG")
	fmt.Println("3. PAKET C/LENGKAP LUAR BIASA")
	fmt.Print("Pilih paket (1/2/3): ")
	var menuChoice string
	fmt.Scanln(&menuChoice)

	switch menuChoice {
	case "1":
		order.Service = "PAKET A/BERSIH AMAN"
	case "2":
		order.Service = "PAKET B/BERSIH TENANG"
	case "3":
		order.Service = "PAKET C/LENGKAP LUAR BIASA"
	default:
		fmt.Println("Pilihan Anda Tidak Ada")
		return
	}

	// Masukkan Unit
	fmt.Print("Masukkan Jumlah Unit (per Kg) :")
	_, err = fmt.Scanln(&order.Unit)
	if err != nil {
		fmt.Println("Input tidak Valid")
		return
	}

	//Pilih Outlet
	fmt.Println("Pilih Outlet:")
	fmt.Println("1. LAUNDRY SENANG")
	fmt.Println("2. LAUNDRY BAHAGIA")
	fmt.Print("Pilih Outlet (1/2/): ")
	var outletChoice string
	fmt.Scanln(&outletChoice)

	switch outletChoice {
	case "1":
		order.OutletName = "LAUNDRY SENANG"
	case "2":
		order.OutletName = "LAUNDRY BAHAGIA"
	default:
		fmt.Println("Pilihan tidak valid")
		return
	}

	// Input Tanggal Order
	fmt.Print("Masukkan Tanggal Order (YYYY-MM-DD): ")

	var orderDateStr string
	_, err = fmt.Scanln(&orderDateStr)
	if err != nil {
		fmt.Println("Input Tanggal Tidak Valid")
	}

	// Validasi format tanggal
	layout := "2006-01-02"
	parseDate, err := time.Parse(layout, orderDateStr)
	if err != nil {
		fmt.Println("Format tanggal tidak valid. Harap Masukkan dengan format yang sudah ditentukan")
		return
	}

	// Set order.OrderDate ke hasil parsing yang valid
	order.OrderDate = parseDate.Format(layout)

	// Pilih Status Order
	fmt.Println("Pilih Status Order :")
	fmt.Println("1. Dalam Proses")
	fmt.Println("2. Selesai")
	fmt.Println("3. Dibatalkan")
	fmt.Print("Pilih status (1/2/3): ")
	var statusChoice string
	fmt.Scanln(&statusChoice)

	switch statusChoice {
	case "1":
		order.Status = "Proses"
	case "2":
		order.Status = "Done"
	case "3":
		order.Status = "Cancel"
	default:
		fmt.Println("Pilihan tidak ada")
		return
	}

	sqlStatement := "INSERT INTO orders (order_id, cust_id, cust_name, service, unit, outlet_name, order_date, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"

	_, err = db.Exec(sqlStatement, order.OrderId, order.CustomerId, order.CustomerName, order.Service, order.Unit, order.OutletName, order.OrderDate, order.Status)
	if err != nil {
		fmt.Println("Gagal Menambahkan Pesanan: ", err)
	} else {
		fmt.Println("Succesfully Added Order!")
	}

}

// update order
func updateOrder(order entity.Orders) {
	db := connectDb()
	defer db.Close()
	var err error

	fmt.Println("=== UBAH DATA TRANSAKSI ===")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Transaksi : ")
	order.OrderId, _ = reader.ReadString('\n')
	order.OrderId = strings.TrimSpace(order.OrderId)

	//validasi ID transaksi
	row := db.QueryRow("SELECT cust_id, cust_name, service, unit, outlet_name, order_date, status FROM orders WHERE order_id = $1;", order.OrderId)
	err = row.Scan(&order.CustomerId, &order.CustomerName, &order.Service, &order.Unit, &order.OutletName, &order.OrderDate, &order.Status)
	if err == sql.ErrNoRows {
		fmt.Println("Transaksi tidak ditemukan!")
		return
	} else if err != nil {
		panic(err)
	}

	fmt.Print("Masukkan Nama : ")
	order.CustomerName, _ = reader.ReadString('\n')
	order.CustomerName = strings.TrimSpace(order.CustomerName)

	//validasi id by name
	customerId, err := getCustomerIdByName(db, order.CustomerName)
	if err != nil {
		fmt.Println("Data Pelanggan tidak ditemukan!")
		return
	}
	order.CustomerId = customerId

	fmt.Println("Pilih Paket:")
	fmt.Println("1. PAKET A/BERSIH AMAN")
	fmt.Println("2. PAKET B/BERSIH TENANG")
	fmt.Println("3. PAKET C/LENGKAP LUAR BIASA")
	fmt.Print("Pilih paket (1/2/3): ")
	var menuChoice string
	fmt.Scanln(&menuChoice)

	switch menuChoice {
	case "1":
		order.Service = "PAKET A/BERSIH AMAN"
	case "2":
		order.Service = "PAKET B/BERSIH TENANG"
	case "3":
		order.Service = "PAKET C/LENGKAP LUAR BIASA"
	default:
		fmt.Println("Pilihan Anda Tidak Ada")
		return
	}

	fmt.Print("Masukkan Unit (dalam Kg): ")
	order.Unit, _ = reader.ReadString('\n')
	order.Unit = strings.TrimSpace(order.Unit)

	fmt.Println("Pilih Outlet:")
	fmt.Println("1. LAUNDRY SENANG")
	fmt.Println("2. LAUNDRY BAHAGIA")
	fmt.Print("Pilih Outlet (1/2/): ")
	var outletChoice string
	fmt.Scanln(&outletChoice)

	switch outletChoice {
	case "1":
		order.OutletName = "LAUNDRY SENANG"
	case "2":
		order.OutletName = "LAUNDRY BAHAGIA"
	default:
		fmt.Println("Pilihan tidak valid")
		return
	}

	fmt.Print("Masukkan Tanggal Order (YYYY-MM-DD): ")
	_, err = fmt.Scanf("%s\n", &order.OrderDate)
	if err != nil {
		fmt.Println("Input Tanggal Tidak Valid")
	}

	fmt.Println("Pilih Status Order :")
	fmt.Println("1. Dalam Proses")
	fmt.Println("2. Selesai")
	fmt.Println("3. Dibatalkan")
	fmt.Print("Pilih status (1/2/3): ")
	var statusChoice string
	fmt.Scanln(&statusChoice)

	switch statusChoice {
	case "1":
		order.Status = "Proses"
	case "2":
		order.Status = "Done"
	case "3":
		order.Status = "Cancel"
	default:
		fmt.Println("Pilihan tidak ada")
		return
	}

	sqlStatement := "UPDATE orders SET cust_id =$2 ,cust_name = $3, service = $4, unit = $5, outlet_name = $6, order_date = $7, status = $8 WHERE order_id= $1;"

	_, err = db.Exec(sqlStatement, order.OrderId, order.CustomerId, order.CustomerName, order.Service, order.Unit, order.OutletName, order.OrderDate, order.Status)
	if err != nil {
		fmt.Println("Gagal Mengupdate Transaksi", err)
	} else {
		fmt.Println("Successfully Update Data!")
	}
}

// delete Order
func deleteOrder(order entity.Orders) {
	db := connectDb()
	defer db.Close()
	var err error

	fmt.Println("=== HAPUS DATA TRANSAKSI ===")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan ID Transaksi : ")
	order.OrderId, _ = reader.ReadString('\n')
	order.OrderId = strings.TrimSpace(order.OrderId)

	sqlStatement := "DELETE FROM orders WHERE order_id = $1;"

	_, err = db.Exec(sqlStatement, order.OrderId)
	if err != nil {
		fmt.Println("Fail to Delete Transactions", err)
	} else {
		fmt.Println("Successfully Delete Data!")
	}
}

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

// search order by ID
func searchOrderBy(order entity.Orders) {
	db := connectDb()
	defer db.Close()

	fmt.Println("=== CARI TRANSAKSI BERDASARKAN ID ===")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Transaksi: ")
	order.OrderId, _ = reader.ReadString('\n')
	order.OrderId = strings.TrimSpace(order.OrderId)

	if !isOrderExists(db, order.OrderId) {
		fmt.Println("Transaksi tidak ditemukan!")
		return
	}

	sqlStatement := "SELECT order_id, cust_id, cust_name, service, unit, outlet_name, order_date, status FROM orders WHERE order_id = $1;"

	row := db.QueryRow(sqlStatement, order.OrderId)

	err := row.Scan(&order.OrderId, &order.CustomerId, &order.CustomerName, &order.Service, &order.Unit, &order.OutletName, &order.OrderDate, &order.Status)

	if err == sql.ErrNoRows {
		fmt.Println("Transaksi tidak ditemukan.")
	} else if err != nil {
		panic(err)
	}

	fmt.Println("Order ID :", order.OrderId)
	fmt.Println("Customer ID :", order.CustomerId)
	fmt.Println("Nama Customer :", order.CustomerName)
	fmt.Println("Service :", order.Service)
	fmt.Println("Unit :", order.Unit)
	fmt.Println("Outlet :", order.OutletName)
	fmt.Println("Tanggal Order:", order.OrderDate)
	fmt.Println("Status :", order.Status)
}

// search All Orders
func SearchAllOrders(order entity.Orders) {
	db := connectDb()
	defer db.Close()

	fmt.Println("=== CARI SEMUA PELANGGAN ===")

	rows, err := db.Query("SELECT order_id, cust_id, cust_name, service, unit, outlet_name, order_date, status FROM orders;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&order.OrderId, &order.CustomerId, &order.CustomerName, &order.Service, &order.Unit, &order.OutletName, &order.OrderDate, &order.Status)
		if err != nil {
			fmt.Println("Tidak ada data:", err)
			continue
		} else {
			fmt.Println("Order ID :", order.OrderId)
			fmt.Println("Customer ID :", order.CustomerId)
			fmt.Println("Nama Customer :", order.CustomerName)
			fmt.Println("Service :", order.Service)
			fmt.Println("Unit :", order.Unit)
			fmt.Println("Outlet :", order.OutletName)
			fmt.Println("Tanggal Order:", order.OrderDate)
			fmt.Println("Status :", order.Status)
			fmt.Println("============================")
		}
	}
}

// Fungsi untuk memvalidasi apakah sebuah email sudah ada dalam tabel customers
func isEmailExists(db *sql.DB, email string) bool {
	var exists bool
	row := db.QueryRow("SELECT EXISTS (SELECT 1 FROM customers WHERE email = $1)", email)
	err := row.Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists
}

// Fungsi untuk memvalidasi apakah sebuah customer ID sudah ada dalam tabel customers
func isCustomerExists(db *sql.DB, customerID string) bool {
	var exists bool
	row := db.QueryRow("SELECT EXISTS (SELECT 1 FROM customers WHERE cust_id = $1)", customerID)
	err := row.Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists
}

// Fungsi untuk memvalidasi apakah sebuah order ID sudah ada dalam tabel orders
func isOrderExists(db *sql.DB, orderID string) bool {
	var exists bool
	row := db.QueryRow("SELECT EXISTS (SELECT 1 FROM orders WHERE order_id = $1)", orderID)
	err := row.Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists
}
