# Latihan Mandiri Pairing - Interface di Golang

## Deskripsi

Latihan ini dirancang untuk melatih penggunaan interface, empty interface, dan type assertion. Kerjakan bersama pasanganmu dan diskusikan logika sebelum coding.

---

## Tujuan

- Membuat dan menggunakan interface dengan method
- Memahami implementasi struct terhadap interface
- Menggunakan `interface{}` untuk menerima tipe apa pun
- Menerapkan type assertion untuk identifikasi data

---

## Instruksi

1. Buat interface `Penarik` yang memiliki method `TarikTunai(jumlah int)`.
2. Buat 2 struct: `ATM` dan `Teller`, masing-masing memiliki saldo dan mengimplementasikan method tersebut.
3. Buat fungsi `transaksi(p Penarik, j int)` yang menerima interface dan memanggil method-nya.
4. Buat fungsi `cetakInfo(data interface{})` yang:
   - Jika `data` adalah string → cetak jumlah karakter
   - Jika `data` adalah int → cetak apakah genap/ganjil
   - Selain itu → tampilkan "Tipe tidak dikenali"
5. Panggil `transaksi()` dan `cetakInfo()` untuk beberapa data.

---

## Contoh Output

```bash
Tarik tunai melalui ATM: 30000
Tarik tunai melalui Teller: 25000
---
Data: Selamat datang!
Jumlah karakter: 16
---
Data: 45
Bilangan ganjil
---
Data: true
Tipe tidak dikenali
```

## Ketentuan
Dikerjakan secara berpasangan.

File bernama main.go, dapat dijalankan tanpa error.

Tambahkan komentar nama peserta di bagian atas file.
