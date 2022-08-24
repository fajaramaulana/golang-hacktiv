package main

import "fmt"

func main() {
	zayan := "zayan"
	billy := "billy"
	rizky := "rizky"
	fajar := "fajar"
	zaka := "zaka"
	hendar := "hendar"
	wahyu := "wahyu"
	tambunan := "tambunan"
	ramadhan := "ramadhan"
	utomo := "utomo"

	names := []*string{&zayan, &billy, &rizky, &fajar, &zaka, &hendar, &wahyu, &tambunan, &ramadhan, &utomo}

	cetaknih := func(list []*string) {
		for _, name := range list {
			fmt.Println(*name)
		}
	}

	cetaknih(names)
}
