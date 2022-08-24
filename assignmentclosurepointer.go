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
	// zayan := "zayan"
	// billy := "billy"
	// rizky := "rizky"
	// fajar := "fajar"
	// zaka := "zaka"
	// hendar := "hendar"
	// wahyu := "wahyu"
	// tambunan := "tambunan"
	// ramadhan := "ramadhan"
	// utomo := "utomo"

	// names := []*string{&zayan, &billy, &rizky, &fajar, &zaka, &hendar, &wahyu, &tambunan, &ramadhan, &utomo}

	// cetaknih := func(list []*string) []string {
	// 	var result []string
	// 	for _, name := range list {
	// 		result = append(result, *name)
	// 	}

	// 	return result
	// }

	// cetaknihye := cetaknih(names)

	// fmt.Println(cetaknihye)
}
