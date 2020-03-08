package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

const errorExitCode = 1

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
		os.Exit(errorExitCode)
	}

	fmt.Println("exact time:", ntpTime)
}
