package main

import (
	"fmt"
	"github.com/smartwalle/backoff/exponential"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(exponential.Default.Backoff(i))
	}
}
