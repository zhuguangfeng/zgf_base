package main

import (
	"fmt"
	"math/rand"
)

type Element struct {
	Id       int
	NoRandom bool
}

func shuffleArrayWithFixedElement(req []Element) {
	for i := len(req) - 1; i > 0; i-- {
		if !req[i].NoRandom {
			j := rand.Intn(i + 1)
			for req[j].NoRandom {
				j = rand.Intn(i + 1)
			}
			req[i], req[j] = req[j], req[i]
		}
	}
}

func main() {
	var sli = []Element{
		{Id: 1, NoRandom: true},
		{Id: 3, NoRandom: false},
		{Id: 5, NoRandom: false},
		{Id: 4, NoRandom: true},
		{Id: 2, NoRandom: false},
	}
	shuffleArrayWithFixedElement(sli)
	fmt.Println(sli)
}
