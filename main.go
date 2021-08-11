package main

import (
	"fmt"

	"github.com/janjoven/wtp-idgen/idgen"
)

func main() {
	init := idgen.New(1, 2)
	id, _ := init.IdGen()
	fmt.Print(id)

}
