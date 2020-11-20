package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	first, last string
}

func main() {
	fmt.Println("Enter input file name (and press Enter):")
	var fileName string
	fmt.Scan(&fileName)
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	var s []Name
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		tokens := strings.Split(line, " ")
		s = append(s, Name{tokens[0], tokens[1]})
	}
	fmt.Println("Printing name structs from slice:")
	for _, n := range s {
		fmt.Printf("%s %s\n", n.first, n.last)
	}
	f.Close()
}
