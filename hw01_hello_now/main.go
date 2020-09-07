package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

const datetimeLayout = "2006-01-02 15:04:05 -0700 MST"

func formatDatetime(t time.Time) string {
	return t.Round(0).Format(datetimeLayout)
}

func main() {
	debug := false
	_, exists := os.LookupEnv("DEBUG")
	if exists {
		debug = true
	}

	ntpHost, exists := os.LookupEnv("NTP_HOST")
	if !exists {
		ntpHost = "0.beevik-ntp.pool.ntp.org"
		if debug {
			fmt.Printf("No ntp host found in env, use default one: %s\n", ntpHost)
		}
	}

	t := time.Now()
	fmt.Printf("current time: %s\n", formatDatetime(t))

	t, err := ntp.Time(ntpHost)
	if err != nil {
		log.Fatalf("Got error %v\n", err)
	}
	fmt.Printf("exact time: %s\n", formatDatetime(t))
}
