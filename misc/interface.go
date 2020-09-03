package main

import "fmt"

type Animal interface {
	Eat()
	Speak()
	Move()
}

type Cow struct{}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func main() {

	var request, str1, str2 string
	aniMap := make(map[string]Animal)

	cow := Cow{}
	bird := Bird{}
	snake := Snake{}

	for {
		fmt.Print(">")
		fmt.Scan(&request, &str1, &str2)

		if request == "newanimal" {

			switch str2 {
			case "cow":
				aniMap[str1] = cow
			case "bird":
				aniMap[str1] = bird
			case "snake":
				aniMap[str1] = snake
			}
			fmt.Println("Created it!")

		} else if request == "query" {

			ani, exists := aniMap[str1]

			if exists {

				switch str2 {
				case "eat":
					ani.Eat()
				case "move":
					ani.Move()
				case "speak":
					ani.Speak()
				}
			} else {
				fmt.Println("Invalid input.")
			}

		} else {
			fmt.Println("Invalid input")
		}
	}
}
