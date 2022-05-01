package main

import (
	"fmt"
	"math"
	"strings"
)

// interface
// type hitung interface {
// 	luas() float64
// 	keliling() float64
// }

// type lingkaran struct {
// 	diameter float64
// }

// func (l lingkaran) JariJari() float64 {
// 	return l.diameter / 2
// }

// func (l lingkaran) luas() float64 {
// 	var luas float64 = math.Pi * math.Pow(l.diameter/2, 2)
// 	return luas
// }

// func (l lingkaran) keliling() float64 {
// 	var keliling float64 = math.Pi * l.diameter
// 	return keliling
// }

// type persegi struct {
// 	sisi float64
// }

// func (p persegi) luas() float64 {
// 	return math.Pow(p.sisi, 2)
// }

// func (p persegi) keliling() float64 {
// 	return p.sisi * 4
// }

func main() {
	// var bangunDatar hitung

	// bangunDatar = persegi{20.5}
	// fmt.Println("Luas Persegi", bangunDatar.luas())
	// fmt.Println("Keliling Persegi", bangunDatar.keliling())

	// bangunDatar = lingkaran{10.5}
	// fmt.Println("Luas Lingkaran", bangunDatar.luas())
	// fmt.Println("Keliling Lingkaran", bangunDatar.keliling())
	// // karena lingkaran tidak terdefinisi didalam interface hitung,
	// // maka pemanggilannya haru dicasting terlebih dahulu
	// fmt.Println("Jari-Jari Lingkaran", bangunDatar.(lingkaran).JariJari())

	var bangunRuang = &kubus{25.39}

	fmt.Println("Luas", bangunRuang.luas())
	fmt.Println("Keliling", bangunRuang.keliling())
	fmt.Println("Volume", bangunRuang.volume())

	// Penggunaan Interface Kosong
	var secret interface{}

	secret = "secret"
	fmt.Println(secret)

	secret = math.Pi
	fmt.Println(secret)

	// interface kosong slice
	secret = []string{"secret", "secret2"}
	fmt.Println(secret)

	secret = map[string]string{"secret": "secret", "secret2": "secret2"}
	fmt.Println(secret)

	var data map[string]interface{}
	data = map[string]interface{}{
		"secret":  "secret",
		"secret2": "secret2",
		"secret3": []string{"secret", "secret2"},
		"secret4": map[string]string{"secret": "secret", "secret2": "secret2"},
	}

	fmt.Println(data["secret4"])

	// casting variable bertipe interface kosong
	var secret1 interface{}
	secret1 = 10
	var number int = secret1.(int) * 10 // casting tipe interface ke integer
	fmt.Println(number)

	secret1 = []string{"secret", "secret2"}
	var secrets = strings.Join(secret1.([]string), " ")
	fmt.Println(secrets)

	// Casting variable interface kosong ke objek pointer
	var orang interface{} = &person{"John", 12}
	var name = orang.(*person).name

	fmt.Println(name)

}

type person struct {
	name string
	age  int
}

// Embeded Interface
type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}
