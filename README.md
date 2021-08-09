# WTP ID Generator


To install this package
```go
 go get github.com/janjoven/wtp-idgen 

```
and import it to your file to use
```go
import (
	"github.com/janjoven/wtp-idgen/idgen"
)
```
IdGenerator Takes to argument in initializing
first is the generation type (1 => RoundId,2 => TransactionId)
return a string

Sample usage
```go
package main

import (
	"fmt"

	"github.com/janjoven/wtp-idgen/idgen"
)

func main() {
	init := idgen.New(
        1, //this is the generation type(1=> RoundId,2 => TransactionID)
        2 //this is the provider Id to specify
        )
	fmt.Print(init.IdGen() + "\n")
}

```
