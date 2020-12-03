package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter up to 10 integers separated by spaces:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tokens := strings.Split(scanner.Text(), " ")
	if (len(tokens) > 10){
       fmt.Println("You entered more than 10 ints :(")
       return
    }
	var numbers = make([]int, len(tokens))
	for i, int_str := range tokens {
		num, _ := strconv.Atoi(int_str)
		numbers[i] = num
	}
	bubbleSort(numbers)
    fmt.Println("Sorted ints:")
	fmt.Println(numbers)
}

func bubbleSort(slice []int) {
	for i := len(slice) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if slice[j] > slice[j+1] {
				swap(slice, j)
			}
		}
	}
}

func swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}
