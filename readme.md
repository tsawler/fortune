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
	"fmt"
	"github.com/tsawler/fortune"
	"log"
	"net/http"
	"time"
)

func main() {
	myFortune := fortune.API{
		Client: &http.Client{Timeout: 10 * time.Second},
		Url:    "https://fortunecookieapi.herokuapp.com/v1/fortunes",
	}

	theFortune, err := myFortune.RandomFortune()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(theFortune)
}

```
