package main

import "fmt"

type BangunData interface {
	HitungLuas() float64
}

type Persegi struct {
	sisi float64
}

type Lingkaran struct {
	jariJari float64
}

func (p *Persegi) HitungLuas() float64 {
	return p.sisi * p.sisi
}

func (l *Lingkaran) HitungLuas() float64 {
	return 3.14 * l.jariJari * l.jariJari
}

func main() {
	persegi := Persegi{sisi: 1.5}
	lingkaran := Lingkaran{jariJari: 5.0}

	luasPersegi := persegi.HitungLuas()
	luasLingkaran := lingkaran.HitungLuas()

	fmt.Println("Luas Persegi:", luasPersegi)
	fmt.Println("Luas Lingkaran:", luasLingkaran)
}