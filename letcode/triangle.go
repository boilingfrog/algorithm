package main

import (
	"bufio"
	"fmt"
	"os"
)

func MakeTriangle() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
	fmt.Printf("bufio.NewScanner:%q\r\n", scanner.Text())

}
