package main

import (
	"fmt"
	"github.com/smartwalle/backoff"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(backoff.Default.Duration(i))
	}
}
