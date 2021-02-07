package main

import (
	"fmt"
	"log"
)

func d() func(string) error {
	return func(p string) error {
		fmt.Println("function d param: ", p)
		return nil
	}
}

func e(p string) {
	fmt.Println("function e parameter:", p)
}

func f(p string) {
	fmt.Println("function f parameter:", p)
}

func g(p string, q int) {
	fmt.Println("function g parameters:", p, q)
}

func main() {
	m := map[string]interface{}{
		"f": f,
		"g": g,
	}

	m["f"].(func(string))("asd")

	s := map[string]func(string){
		"e": e,
		"f": f,
	}

	cmd, ok := s["e"]
	if !ok {
		log.Fatal("key now found")
	}
	cmd("Asd")

	r := map[string]func() func(string) error{
		"d": d,
	}

	cmd2, ok := r["d"]
	if !ok {
		log.Fatal("key now found")
	}
	cmd2()("asdfasdasdad")
}
