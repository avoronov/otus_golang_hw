package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	t := time.Now()
	fmt.Printf("current time: %v\n", t.Round(0))

	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("Got error %v\n", err)
	}
	fmt.Printf("exact time: %v\n", t.Round(0))
}
