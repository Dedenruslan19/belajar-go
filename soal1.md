# Latihan Mandiri Slice - Golang

## Konteks
Tugas ini dirancang sebagai latihan mandiri berpasangan (_pairing assignment_) untuk menguatkan pemahaman konsep **slice di Golang**, dengan konteks **aplikasi keuangan sederhana**. Latihan ini akan melatih kalian dalam manipulasi data dinamis, fungsi, dan pengolahan slice secara efisien.

> 💡 Format pengerjaan: **1 pasang 2 orang**. Diskusikan logika bersama, lalu implementasikan solusinya di 1 file Go.

---

## Tujuan Pembelajaran
- Menggunakan slice untuk menyimpan data dinamis.
- Menambahkan, menghapus, dan memfilter data dalam slice.
- Menggunakan slice sebagai parameter fungsi.
- Menerapkan logika kontrol dalam konteks real-life (saldo nasabah).

---

## Requirements

### 1. Menyimpan dan Menambahkan Saldo
- Buat slice dengan 8 data saldo awal nasabah (bebas nilainya).
- Tambahkan 2 saldo baru menggunakan `append()`.

> ⛏️ *Hint:* Gunakan slice literal dan `append`.

---

### 2. Fungsi Filter Saldo > 25.000
- Buat fungsi `cetakSaldoTinggi(s []int)` yang mencetak semua saldo nasabah **yang lebih dari 25.000**.
- Fungsi ini tidak mengembalikan nilai, hanya mencetak data yang sesuai.

> ⛏️ *Hint:* Gunakan `for` dan `if`.

---

### 3. Fungsi Hapus Saldo Pertama
- Buat fungsi `hapusSaldoPertama(s []int) []int` yang menghapus elemen pertama dari slice dan **mengembalikan slice baru**.
- Tampilkan hasil slice setelah fungsi dipanggil.

> ⛏️ *Hint:* Gunakan slicing dan `append`.

---

## Contoh Struktur Output Program (Opsional)

```bash
Saldo awal: [10000 20000 15000 40000 50000 32000 18000 27000]
Saldo setelah penambahan: [10000 20000 15000 40000 50000 32000 18000 27000 30000 35000]
Saldo > 25000:
40000
50000
32000
27000
30000
35000
Saldo setelah hapus pertama: [20000 15000 40000 50000 32000 18000 27000 30000 35000]
```

## Ketentuan Tugas
Submit 1 file Go (main.go) dari hasil kerja berpasangan.

Sertakan komentar nama kedua peserta di bagian atas file.

Jangan hanya membagi tugas penulisan, pastikan kalian berdua berdiskusi dan memahami logikanya.

Deadline: [tentukan oleh instruktur]

## Catatan Tambahan
Kode harus dapat dijalankan tanpa error.

Gunakan komentar secukupnya untuk menjelaskan alur logika.

