package main

import (
	"fmt"
	"os"
	"strconv"
)

type Orang struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	orang := []Orang{
		{
			Nama:      "Fajar",
			Alamat:    "Jl. Kebon Jeruk",
			Pekerjaan: "apa hayo",
			Alasan:    "Biar Bisa Golang",
		},
		{
			Nama:      "Zaka",
			Alamat:    "Jl. Kebon Jeruk",
			Pekerjaan: "apa hayo",
			Alasan:    "Biar Keren Aja",
		},
		{
			Nama:      "Zayan",
			Alamat:    "Jl. Kebon Jeruk",
			Pekerjaan: "apa hayo",
			Alasan:    "Biar Bisa Golang",
		},
	}

	fmt.Println(nyetakOrang(orang))
}

func nyetakOrang(list []Orang) Orang {
	// var result []string
	num, _ := strconv.Atoi(os.Args[1])
	return list[num]
}
