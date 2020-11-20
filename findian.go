package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter string (and press Enter):")
	match := true
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x := strings.ToLower(scanner.Text())
	if x[0] != 'i' {
		match = false
	}
	if x[len(x)-1] != 'n' {
		match = false
	}
	middle_a := false
	for i := 1; i < len(x)-1; i++ {
		if x[i] == 'a' {
			middle_a = true
		}
	}
	if middle_a == false {
		match = false
	}
	if match == true {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
