package main

import "fmt"

type Penggajian interface {
	HitungGaji() float64
}

type KaryawanTetap struct {
	nama        string
	gajiBulanan float64
}

func (k KaryawanTetap) HitungGaji() float64 {
	return k.gajiBulanan
}

type KaryawanKontrak struct {
	nama       string
	gajiPerJam float64
	jamKerja   float64
}

func (k KaryawanKontrak) HitungGaji() float64 {
	return k.gajiPerJam * k.jamKerja
}

func cetakSlipGaji(p Penggajian) {
	switch v := p.(type) {
	case KaryawanTetap:
		fmt.Printf("Slip gaji: %s \n", v.nama)
		fmt.Printf("Total gaji: %.2f \n", v.HitungGaji())
	case KaryawanKontrak:
		fmt.Printf("Slip gaji: %s \n", v.nama)
		fmt.Printf("Total gaji: %.2f \n", v.HitungGaji())
	default:
		fmt.Println("Type karyawan tidak dikenal.")
	}
}

func main() {
	// Buat objek KaryawanTetap
	karyawan1 := KaryawanTetap{
		nama:        "Budi",
		gajiBulanan: 8000000.0,
	}

	// Buat objek KaryawanKontrak
	karyawan2 := KaryawanKontrak{
		nama:        "Citra",
		gajiPerJam:  50000.0,
		jamKerja:    160.0,
	}

	// Cetak slip gaji untuk kedua karyawan menggunakan satu fungsi
	cetakSlipGaji(karyawan1)
	cetakSlipGaji(karyawan2)
}