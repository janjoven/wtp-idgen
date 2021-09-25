package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/janjoven/wtp-idgen/idgen"
)

var prevNumber int64
var mu sync.Mutex

func main() {
	init := idgen.New(1000000, 2)
	for {
		mu.Lock()
		RoundID, err := init.IdGenTransaction(42, 1)
		if err != nil {
			log.Printf("%v", err)
		}
		if prevNumber == RoundID {
			break
		}
		fmt.Println(RoundID)
		prevNumber = RoundID
		mu.Unlock()
		time.Sleep(10 * time.Millisecond)
	}
}
