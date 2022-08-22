package main

import "fmt"

func main() {
	names := []string{"zayan", "billy", "rizky", "fajar", "zaka", "hendar", "wahyu", "tambunan", "ramadhan", "utomo"}

	for _, name := range names {
		cetak(name)
	}

}

func cetak(name string) {
	fmt.Println(name)
}
