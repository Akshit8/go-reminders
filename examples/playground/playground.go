package main

import "fmt"

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
		"f": f,
	}

	s["f"]("asd2")
}
