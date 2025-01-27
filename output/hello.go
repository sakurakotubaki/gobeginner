package output

type Greet struct {
	Value string
}

func (g Greet) Hello() string {
	return g.Value
}

/*
package main

import "gobeginner/output"

func main() {
	g := output.Greet{Value: "hi gopher"}
	println(g.Hello())
}
*/
