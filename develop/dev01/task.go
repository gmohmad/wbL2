package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Basic task ===

Create a program that prints the exact time using the NTP library. Initialize it as a go module.
Use the library https://github.com/beevik/ntp.
Write a program that prints the current time / exact time using this library.

The program must be formatted using a go module.
The program must correctly handle library errors: print them to STDERR and return a non-zero exit code to the OS.
The program must pass the go vet and golint checks.
*/

func main() {
	// Get exact current time using ntp.Time
	currTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	// If error occurs, write the error to Stderr and exit with code 1
	if err != nil {
		io.WriteString(os.Stderr, err.Error())
		fmt.Println()
		os.Exit(1)
	}

	// Else print current time using time.Now and the exact time we got using ntp.Time
	fmt.Println(time.Now())
	fmt.Println(currTime)
}
