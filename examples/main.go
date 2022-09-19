package main

import (
	"fmt"
	"github.com/smartwalle/backoff/exponential"
)

func main() {
	var backoff = exponential.New()

	for i := 0; i < 10; i++ {
		fmt.Println(backoff.Backoff(i))
	}
}
