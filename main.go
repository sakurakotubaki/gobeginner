package main

import "gobeginner/output"

func main() {
	g := output.Greet{Value: "hi gopher"}
	println(g.Hello())
}
