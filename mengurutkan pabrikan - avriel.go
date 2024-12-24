package main

import (
	"fmt"
	"sort"
)

// Tipe bentukan untuk data mobil
type Mobil struct {
	ID_mobil   int
	NamaMobil  string
	IDpabrikan int
	Tahun      int
	HargaMobil float64
}

// Tipe bentukan untuk data pabrikan
type Pabrikan struct {
	ID          int
	Nama        string
	Wilayah     string
	HitungMobil int
}

// Deklarasi array statis untuk menyimpan data mobil dan pabrikan
const maxCars = 100

var mobilList [maxCars]Mobil

const maxPabrikan = 100

var pabrikanList [maxPabrikan]Pabrikan

var mobilCount int = 0
var pabrikanCount int = 0

// Fungsi untuk mengurutkan nama pabrikan berdasarkan jumlah mobil yang ada
func sortPabrikanByCarCount(ascending bool) {
	sort.SliceStable(pabrikanList[:pabrikanCount], func(i, j int) bool {
		if ascending {
			return pabrikanList[i].HitungMobil < pabrikanList[j].HitungMobil
		}
		return pabrikanList[i].HitungMobil > pabrikanList[j].HitungMobil
	})
}

// Fungsi untuk menampilkan mobil terurut berdasarkan tahun keluar, nama, pabrikan, dan lainnya
func displaySortedCars() {
	sort.SliceStable(mobilList[:mobilCount], func(i, j int) bool {
		return mobilList[i].Tahun < mobilList[j].Tahun
	})

	fmt.Println("Daftar mobil terurut:")
	for _, mobil := range mobilList[:mobilCount] {
		fmt.Printf("Nama: %s, Pabrikan ID: %d, Tahun: %d, Harga: %.2f\n",
			mobil.NamaMobil, mobil.IDpabrikan, mobil.Tahun, mobil.HargaMobil)
	}
}

func main() {
	// Urutkan pabrikan berdasarkan jumlah mobil (ascending)
	sortPabrikanByCarCount(true)

	// Tampilkan pabrikan yang terurut berdasarkan jumlah mobil
	fmt.Println("Pabrikan terurut berdasarkan jumlah mobil:")
	for _, pabrikan := range pabrikanList[:pabrikanCount] {
		fmt.Printf("ID: %d, Nama: %s, Jumlah Mobil: %d\n", pabrikan.ID, pabrikan.Nama, pabrikan.HitungMobil)
	}

	// Tampilkan mobil terurut berdasarkan tahun
	displaySortedCars()
}
