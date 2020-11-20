package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]string)
	fmt.Println("Enter name (and press Enter):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input1 := scanner.Text()
	fmt.Println("Enter address (and press Enter):")
	scanner.Scan()
	input2 := scanner.Text()
	m["name"] = input1
	m["address"] = input2
	barr, _ := json.Marshal(m)
	fmt.Println("JSON object:")
	fmt.Println(string(barr))
}
