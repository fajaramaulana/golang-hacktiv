package main

import (
	"fmt"
	"go-belajar-interface/service"
	"sync"
)

func main() {
	var db []*service.User

	userService := service.NewUserService(db)

	res := userService.Register(&service.User{Id: 1, Nama: "Fajar"})
	fmt.Println(res)
	res2 := userService.Register(&service.User{Id: 2, Nama: "Budi"})
	fmt.Println(res2)
	res3 := userService.Register(&service.User{Id: 2, Nama: "Pakong"})
	fmt.Println(res3)

	fmt.Println("-----------Hasil get user------------")
	getAll := userService.GetUser()

	var wg sync.WaitGroup
	wg.Add(len(getAll))
	for _, v := range getAll {
		go printNama(&wg, v)
	}

	wg.Wait()

	// getUserById := userService.GetUserById(2)

	// fmt.Println(getUserById)
}

func printNama(wg *sync.WaitGroup, u *service.User) {
	fmt.Println(u)
	wg.Done()
}
