package main

import (
	"fmt"

	"github.com/janjoven/go-wtp-idgen/idgen"
)

func main() {
	init := idgen.New(1, 2)
	fmt.Print(init.IdGen() + "\n")

}
