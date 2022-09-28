package main

import (
	"fmt"
	"math/rand"
)

type korek struct {
	hits       int
	lastPlayer string
}

func main() {
	c := make(chan *korek)
	done := make(chan *korek)

	go player("User 1", c, done)
	go player("User 2", c, done)
	go player("User 3", c, done)
	go player("User 4", c, done)

	clear(c, done)
}

func player(name string, c, done chan *korek) {
	min := 1
	max := 100
	fmt.Println(rand.Intn(max-min) + min)
}

func clear(c, done chan *korek) {
	k := &korek{}
	for {
		c <- k
		k = <-done
		select {
		case s := <-done:
			if s.hits%11 == 0 {
				fmt.Println("korek ada di", k.lastPlayer, "pada hits", k.hits)
				return
			}
		}
	}
}
