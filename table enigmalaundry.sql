CREATE TABLE customers (
	cust_id VARCHAR(5) PRIMARY KEY, 
	cust_name VARCHAR(100), 
	phone_number VARCHAR(20),
	address VARCHAR (100),
	email VARCHAR (100)
);

ALTER TABLE customers 
ADD CONSTRAINT cust_id_unique UNIQUE (cust_id);
ALTER TABLE customers 
ADD CONSTRAINT email_unique UNIQUE (email);


CREATE TABLE services (
	service_id VARCHAR(5) PRIMARY KEY,
	service_name VARCHAR(100),
	unit INT,
	price DECIMAL(10, 2)
);

ALTER TABLE services
ADD CONSTRAINT service_id_unique UNIQUE (service_id);

INSERT INTO services (
	service_id,
	service_name,
	unit,
	price
) VALUES 
	('SP001', 'PAKET A/BERSIH AMAN', '1', '7000'),
	('SP002', 'PAKET B/BERSIH TENANG', '1', '8000'),
	('SP003', 'PAKET C/LENGKAP LUAR BIASA', '1', '12000');
	


CREATE TABLE outlets (
	outlet_id VARCHAR(5) PRIMARY KEY,
	outlet_name VARCHAR(100),
	outlet_address VARCHAR(100)
);

INSERT INTO outlets (
	outlet_id,
	outlet_name,
	outlet_address
) VALUES 
	('OT001', 'LAUNDRY SENANG', 'TANGERANG'),
	('OT002', 'LAUNDRY BAHAGIA', 'JAKARTA BARAT');


ALTER TABLE outlets
ADD CONSTRAINT outlet_id_unique UNIQUE (outlet_id);


CREATE TABLE orders (
	order_id VARCHAR(5) PRIMARY KEY,
	cust_id VARCHAR(5) REFERENCES customers(cust_id),
	cust_name VARCHAR(100),
	service VARCHAR(100),
	unit INT,
	outlet_name VARCHAR(100),
	order_date DATE,
	status VARCHAR(10)
);

ALTER TABLE orders 
ADD CONSTRAINT order_id_unique UNIQUE (order_id);