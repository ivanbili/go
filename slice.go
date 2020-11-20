package main

import "fmt"
import "strconv"

func main() {
	arr := make([]int, 0, 3)
	for {
		fmt.Println("Enter number (or X to quit):")
		var str string
		fmt.Scan(&str)
		if str == "X" {
			break
		}
		num, _ := strconv.Atoi(str)
		arr = append(arr, num)
		fmt.Println("Sorted numbers:")
		for i := len(arr) - 1; i > 0; i-- {
			if arr[i] > arr[i-1] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
			fmt.Printf("%d ", arr[i])
		}
		fmt.Printf("%d\n", arr[0])
	}
}
