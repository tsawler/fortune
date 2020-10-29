# Fortune Cookie Generator

This is just a sample go module that returns a random fortune.

Installation:

```
go get -u github.com/tsawler/fortune
```



Usage:

```go
package main

import (
	"github.com/tsawler/fortune"
	"log"
)

func main() {
	myFortune, err := fortune.RandomFortune()
	if err != nil {
		log.Println(err)
	}
	log.Println(myFortune)
}
```
