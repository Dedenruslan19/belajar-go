# Latihan Mandiri Map - Golang

## Deskripsi Tugas

Latihan ini dirancang untuk dikerjakan secara berpasangan (pairing) oleh dua peserta. Tujuannya adalah untuk memperkuat pemahaman tentang penggunaan map di Golang dalam konteks aplikasi keuangan sederhana.

Silakan diskusikan logika dan implementasi bersama pasangan kalian, dan kerjakan dalam satu file Go (`main.go`).

## Tujuan Pembelajaran

- Menggunakan map untuk menyimpan data dengan key-value.
- Menambahkan, mengubah, dan menghapus entri dalam map.
- Mengakses dan memfilter data berdasarkan kondisi tertentu.
- Membuat fungsi dengan parameter bertipe map.

## Instruksi

1. Buat map saldo nasabah dengan minimal 5 entri.
   - Tipe: `map[string]int`, dengan nama nasabah sebagai key dan saldo sebagai value.
   - Contoh: `"Andi": 15000`

2. Tambahkan satu nasabah baru bernama `"Agus"` dengan saldo `30000`.

3. Buat fungsi dengan nama `hitungSaldoTinggi`:
   - Parameter: map nasabah dan batas saldo (integer).
   - Fungsi ini harus mengembalikan slice berisi nama-nama nasabah yang memiliki saldo lebih besar dari nilai batas tersebut.
   - Tipe kembalian: `[]string`

4. Hapus satu entri nasabah dari map yang kalian buat di awal.

5. Cetak isi map setelah penghapusan, dan tampilkan hasil dari fungsi `hitungSaldoTinggi`.

# Contoh Output yang Diharapkan

### Map awal: map[Andi:15000 Budi:18000 Citra:27000 Dini:30000 Eko:12000]

### Map setelah tambah Agus: map[Agus:30000 Andi:15000 Budi:18000 Citra:27000 Dini:30000 Eko:12000]

### Map setelah hapus Budi: map[Agus:30000 Andi:15000 Citra:27000 Dini:30000 Eko:12000]

#### Nasabah dengan saldo di atas 25000:  
Citra  
Dini  
Agus


## Ketentuan Tugas

- Kode harus bebas error dan dapat dijalankan.
- Tulis nama kedua peserta di bagian atas file Go sebagai komentar.
- Boleh menggunakan `range`, `append`, dan `delete`.
- Tidak diperbolehkan menggunakan slice untuk menyimpan semua data utama, hanya untuk hasil filter (jika dibutuhkan).
