package main

import "fmt"

type Manusia struct {
	nama string
}

func main() {

	names := []*Manusia{{nama: "fajar"}, {nama: "zaka"}, {nama: "hendar"}, {nama: "wahyu"}, {nama: "tambunan"}, {nama: "ramadhan"}, {nama: "utomo"}}

	cetaknih := func(list []*Manusia) []string {
		// var result []string
		// for _, name := range list {
		// 	result = append(result, *name)
		// }

		// return result

		var result []string
		for _, name := range list {
			result = append(result, name.nama)
		}

		return result
	}

	fmt.Println(cetaknih(names))
}
