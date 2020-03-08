package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	printLocalTime()
	printNtpTime()
}

func printLocalTime() {
	fmt.Println("current time:", time.Now())
}

func printNtpTime() {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println("exact time:", ntpTime)
}
