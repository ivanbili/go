package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println("Enter integers separated by spaces:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tokens := strings.Split(scanner.Text(), " ")
	var numbers = make([]int, len(tokens))
	for i, int_str := range tokens {
		num, _ := strconv.Atoi(int_str)
		numbers[i] = num
	}
	half1 := numbers[:len(numbers)/2]
	half2 := numbers[len(numbers)/2:]
	q1 := half1[:len(half1)/2]
	q2 := half1[len(half1)/2:]
	q3 := half2[:len(half2)/2]
	q4 := half2[len(half2)/2:]
	fmt.Println("here are the initial 4 slices:")
	fmt.Println(q1)
	fmt.Println(q2)
	fmt.Println(q3)
	fmt.Println(q4)
	fmt.Println("mergeSort called concurrently on the following slices:")
	var wg sync.WaitGroup
	wg.Add(4)
	go mergeSort(q1, &wg)
	go mergeSort(q2, &wg)
	go mergeSort(q3, &wg)
	go mergeSort(q4, &wg)
	wg.Wait()
	wg.Add(2)
	go merge(q1, q2, &wg)
	go merge(q3, q4, &wg)
	wg.Wait()
	merge(half1, half2, nil)
	fmt.Println("Sorted ints:")
	fmt.Println(numbers)
}

func mergeSort(slice []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(slice)
	if len(slice) <= 1 {
		return
	}
	h1 := slice[:len(slice)/2]
	h2 := slice[len(slice)/2:]
	var wgInner sync.WaitGroup
	wgInner.Add(2)
	go mergeSort(h1, &wgInner)
	go mergeSort(h2, &wgInner)
	wgInner.Wait()
	merge(h1, h2, nil)
}

func merge(slice1 []int, slice2 []int, wg *sync.WaitGroup) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return
	}
	result := make([]int, 0, len(slice1)+len(slice2))
	var left int = 0
	var right int = 0
	for i := 0; i < len(slice1)+len(slice2); i++ {
		if right == len(slice2) || (left != len(slice1) && slice1[left] < slice2[right]) {
			result = append(result, slice1[left])
			left++
		} else {
			result = append(result, slice2[right])
			right++
		}
	}
	copy(slice1, result)
	copy(slice2, result[len(slice1):])
	if wg != nil {
		wg.Done()
	}
}
