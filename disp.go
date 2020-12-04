package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter values for acceleration (a), initial velocity (v0), and initial displacement (s0), separated by spaces:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tokens := strings.Split(scanner.Text(), " ")
	if (len(tokens) != 3){
       fmt.Println("You didn't enter 3 numbers :(")
       return
    }
	var numbers = make([]float64, len(tokens))
	for i, float_str := range tokens {
		num, _ := strconv.ParseFloat(float_str, 64);
		numbers[i] = num
	}
	fn := GenDisplaceFn(numbers[0], numbers[1], numbers[2])
	fmt.Println("Enter time: ")
    scanner.Scan()
    time, _ := strconv.ParseFloat(scanner.Text(), 64);
    fmt.Printf("Position at time t = %f is:\n", time)
	fmt.Println(fn(time))
}

func GenDisplaceFn(a, v0, s0 float64) func (float64) float64{
     fn := func (t float64) float64 {
         return 0.5*a*t*t + v0*t + s0
     }
     return fn
}
