package main

import "fmt"

func main() {
	fmt.Println("Enter floating point number (and press Enter):")
	var x float64
	fmt.Scan(&x)
	fmt.Println("Truncated number is:", int(x))
}
