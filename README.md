# enigma-laundry-project

Final Project from Enigma Camp course Golang DB

## Instalasi

Clone repositori dari GitLab ke mesin lokal Anda:

```bash
git clone https://gitlab.com/username-anda/enigma-laundry.git
```

Buat folder enigma-laundry lalu beralih ke folder tersebut dengan menggunaakn git bash/CMD/PowerShell
cd enigma-laundry

Pasang dependensi yang diperlukan. Pastikan Anda telah memiliki Go pada Komputer/Laptop anda.

go get -u github.com/lib/pq

Siapkan database PostgreSQL:

Pastikan Anda telah memiliki PostgreSQL terpasang dan berjalan.
Buat database dengan nama 'enigmalaundry' dengan pengaturan yang sesuai. Anda dapat mengubah pengaturan koneksi dalam kode jika diperlukan (const pada main.go).

buat tabel pada database 'enigmalaundry' dengan query DML yang ada pada 'table enigmalaundry.sql'

## Penggunaan

go run main.go

Anda akan melihat menu konsol dengan berbagai opsi untuk mengelola pelanggan dan pesanan:
=== Enigma Laundry ===

1. Tambah Pelanggan
2. Ubah Data Pelanggan
3. Hapus Data Pelanggan
4. Tambah Pesanan
5. Ubah Data Pesanan
6. Hapus Data Pesanan
7. Cari Pelanggan
8. Cari Pesanan dan Status
9. Keluar
   Pilih tindakan (1/2/3/4/5/6/7/8/9):

Ikuti petunjuk untuk berinteraksi dengan aplikasi:

Untuk menambahkan pelanggan, pilih opsi 1 dan berikan informasi yang diperlukan.
Untuk mengubah data pelanggan, pilih opsi 2 dan ikuti instruksinya.
Untuk menghapus pelanggan, pilih opsi 3 dan berikan ID pelanggan.
Untuk menambahkan pesanan, pilih opsi 4 dan berikan informasi yang diperlukan.
Untuk mengubah pesanan, pilih opsi 5 dan ikuti instruksinya.
Untuk menghapus pesanan, pilih opsi 6 dan berikan ID pesanan.
Untuk mencari pelanggan atau pesanan, pilih opsi yang sesuai dan ikuti petunjuknya.
Pilih opsi 9 untuk keluar dari aplikasi.

## CATATAN

- Pastikan untuk mengonfigurasi pengaturan koneksi PostgreSQL dalam berkas main.go sebelum menjalankan aplikasi.

- cust_id pada tabel customers dan order_id pada tabel orders adalah auto increment dan unique, jadi tidak perlu di input manual oleh user

- Pada tabel services anda bisa menambahkan Pilihan Paket, namun jangan lupa untuk menambahkan pilihan tersebut pada fungsi addOrder() dan updateOrder()

-README ini mengasumsikan bahwa Anda telah memiliki Go dan PostgreSQL terpasang serta dikonfigurasi pada mesin Anda.

-Aplikasi ini di posting sebagai bagian yang bertujuan untuk final projek pada course Golang DB Enigma Camp dan mungkin akan author gunakan untuk kepentingan portofolio dan sebagai prototype aplikasi yang akan dibuat di masa mendatang.

- Karena keterbatasan waktu dan juga author yang sambil bekerja mengikuti bootcamp ini maka aplikasi ini dibuat dengan sangat sederhana, mungkin kedepannya akan di update sebagai bahan portofolio author ataupun aplikasi yang akan di produksi

Untuk masalah atau pertanyaan, silakan hubungi saya di discord xrmzy / Ramzi#0440.
