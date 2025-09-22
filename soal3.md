Soal Latihan
Tugas: Buatlah sebuah program untuk menghitung dan mencetak biaya pengiriman barang ke berbagai lokasi. Biaya pengiriman ditentukan oleh dua hal: jarak dan berat.

Instruksi:

Buatlah sebuah interface bernama Pengiriman yang memiliki satu method: HitungBiaya() float64.

Buatlah sebuah struct bernama PengirimanKilat yang memiliki properti berat (float64) dan jarak (float64).

Buatlah sebuah struct lain bernama PengirimanStandar yang juga memiliki properti berat (float64) dan jarak (float64).

Implementasikan method HitungBiaya() untuk kedua struct tersebut:

Untuk PengirimanKilat, biaya dihitung dengan rumus: (berat * 15000) + (jarak * 1000).

Untuk PengirimanStandar, biaya dihitung dengan rumus: (berat * 10000) + (jarak * 500).

Di fungsi main(), buatlah sebuah objek PengirimanKilat dan satu objek PengirimanStandar. Berikan nilai untuk berat dan jarak pada kedua objek tersebut.

Gunakan satu fungsi (misalnya, cetakBiaya(p Pengiriman)) yang menerima interface Pengiriman untuk mencetak biaya total pengiriman dari kedua objek tersebut.