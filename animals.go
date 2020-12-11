package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var cow Animal
	cow.Init("grass", "walk", "moo")
	var bird Animal
	bird.Init("worms", "fly", "peep")
	var snake Animal
	snake.Init("mice", "slither", "hsss")
	fmt.Println("Enter animal (“cow”, “bird”, or “snake”) and command(“eat”, “move”, or “speak”):")
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		tokens := strings.Split(scanner.Text(), " ")
		var a Animal
		if tokens[0] == "cow" {
			a = cow
		} else if tokens[0] == "bird" {
			a = bird
		} else if tokens[0] == "snake" {
			a = snake
		} else {
			fmt.Println("Invalid animal :( ")
			continue
		}
		if tokens[1] == "eat" {
			a.Eat()
		} else if tokens[1] == "move" {
			a.Move()
		} else if tokens[1] == "speak" {
			a.Speak()
		} else {
			fmt.Println("Invalid command :( ")
		}
	}
}

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Init(f, l, n string) {
	a.food = f
	a.locomotion = l
	a.noise = n
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}
