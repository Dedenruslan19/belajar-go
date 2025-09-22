Soal Latihan
Tugas: Buatlah program untuk mengelola data karyawan di sebuah perusahaan.

Instruksi:

Buatlah sebuah interface bernama Penggajian yang memiliki satu method: HitungGaji() float64.

Buatlah sebuah struct bernama KaryawanTetap yang memiliki properti nama (string) dan gajiBulanan (float64).

Buatlah sebuah struct lain bernama KaryawanKontrak yang memiliki properti nama (string), gajiPerJam (float64), dan jamKerja (float64).

Implementasikan method HitungGaji() untuk kedua struct tersebut.

Untuk KaryawanTetap, method ini cukup mengembalikan nilai gajiBulanan.

Untuk KaryawanKontrak, method ini harus mengalikan gajiPerJam dengan jamKerja untuk mendapatkan total gaji.

Buat sebuah fungsi bernama cetakSlipGaji(p Penggajian) yang menerima interface Penggajian dan mencetak nama serta total gaji.

Di fungsi main(), buatlah sebuah objek KaryawanTetap dan satu lagi KaryawanKontrak.

Panggil fungsi cetakSlipGaji() untuk kedua objek tersebut untuk menampilkan gaji masing-masing.