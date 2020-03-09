package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
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
		log.Fatalln(err.Error())
	}

	fmt.Println("exact time:", ntpTime)
}
