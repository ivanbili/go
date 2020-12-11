package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var farm []NamedEntity
	fmt.Println("Enter command ('newanimal' or 'query')")
	fmt.Println("e.g. 'newanimal ivan cow'")
	fmt.Println("     'query ivan eat'")
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		tokens := strings.Split(scanner.Text(), " ")

		if tokens[0] == "newanimal" {
			var ne NamedEntity
			if tokens[2] == "cow" {
				ne = &Cow{tokens[1]}
			} else if tokens[2] == "bird" {
				ne = &Bird{tokens[1]}
			} else if tokens[2] == "snake" {
				ne = &Snake{tokens[1]}
			}
			farm = append(farm, ne)
			fmt.Println("Created it!")
		} else if tokens[0] == "query" {

			var ne NamedEntity
			for _, cur := range farm {
				if cur.Name() == tokens[1] {
					ne = cur
					break
				}
			}
			if ne == nil {
				fmt.Println("Couldn't find that name :(")
				continue
			}
			var a, ok = ne.(Animal)
			if ok {
				if tokens[2] == "eat" {
					a.Eat()
				} else if tokens[2] == "move" {
					a.Move()
				} else if tokens[2] == "speak" {
					a.Speak()
				} else {
					fmt.Println("Don't know this action :(")
				}
			} else {
				fmt.Println("This is not an animal :(")
			}
		} else {
			fmt.Println("Invalid command :( ")
			continue
		}
	}
}

type NamedEntity interface {
	Name() string
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }

func (a *Cow) Eat() {
	fmt.Println("grass")
}

func (a *Cow) Move() {
	fmt.Println("walk")
}

func (a *Cow) Speak() {
	fmt.Println("moo")
}

func (a *Cow) Name() string {
	return a.name
}

type Bird struct{ name string }

func (a *Bird) Eat() {
	fmt.Println("worms")
}

func (a *Bird) Move() {
	fmt.Println("fly")
}

func (a *Bird) Speak() {
	fmt.Println("peep")
}

func (a *Bird) Name() string {
	return a.name
}

type Snake struct{ name string }

func (a *Snake) Eat() {
	fmt.Println("mice")
}

func (a *Snake) Move() {
	fmt.Println("slither")
}

func (a *Snake) Speak() {
	fmt.Println("hsss")
}

func (a *Snake) Name() string {
	return a.name
}
