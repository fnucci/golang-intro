package main

import (
	"fmt"
	"time"
)

func main() {
	// data := "10/Jun/2019 - 08:30:45"

	data := time.Now()
	fmt.Println(data.Format("02-01-2006 03:04:05"))
	fmt.Println(data.Format("Mon, 02 de Jan de 2006"))
}
