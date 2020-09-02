package main

import "fmt"

type Animal struct {
	food, locomotion, noise string
}

func (am Animal) Eat() {
	fmt.Println(am.food)
}

func (am Animal) Move() {
	fmt.Println(am.locomotion)
}

func (am Animal) Speak() {
	fmt.Println(am.noise)
}

func main() {

	an := map[string]Animal{
		"cow" : Animal{"grass","walk","moo"},
		"bird" : Animal{"worms","fly","peep"},
		"snake" : Animal{"mice","slither","hsss"},
	}
	
	for{
		var animalName, animalInfo string
		fmt.Print(">")
		fmt.Scan(&animalName,&animalInfo)
		
		switch animalInfo {
			case "eat":
				an[animalName].Eat()
			case "move":
				an[animalName].Move()
			case "speak":
				an[animalName].Speak()
		}
	}
}
