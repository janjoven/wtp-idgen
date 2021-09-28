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
	for {
		mu.Lock()
		RoundID, err := idgen.IdGenTransaction(42)
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
