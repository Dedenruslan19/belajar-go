package main

import "fmt"

type Pengiriman interface {
	HitungBiaya() float64
}

func (p PengirimanKilat) HitungBiaya() float64 {
	return p.berat * p.jarak * 1000
}

func (p PengirimanStandard) HitungBiaya() float64 {
	return p.berat * p.jarak * 1000
}

type PengirimanKilat struct {
	berat float64
	jarak float64
}

type PengirimanStandard struct {
	berat float64
	jarak float64
}

func cetakBiaya(n Pengiriman) {
	switch v := n.(type) {
	case PengirimanKilat:
		fmt.Printf("Total berat: %2.f\n", v.berat)
		fmt.Printf("Total jarak: %2.f\n", v.jarak)
		fmt.Printf("Harga: %2.f\n",v.HitungBiaya())
	case PengirimanStandard:
		fmt.Printf("Total berat: %2.f\n", v.berat)
		fmt.Printf("Total jarak: %2.f\n", v.jarak)
		fmt.Printf("Harga: %2.f\n",v.HitungBiaya())
	default:
		fmt.Println("Pengiriman tidak diketahui.")
	}
}

func main(){
	pengiriman1 := PengirimanKilat{
		berat : 15,
		jarak : 10,
	}	
	pengiriman2 := PengirimanStandard{
		berat : 15,
		jarak : 2,
	}

	cetakBiaya(pengiriman1)
	cetakBiaya(pengiriman2)
}