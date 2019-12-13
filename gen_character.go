package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type abilities struct {
	strength     int
	dexterity    int
	constitution int
	intelligence int
	wisdom       int
	charisma     int
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func rollDice(d int) int {
	rand.Seed(time.Now().UnixNano())
	randomNum := random(1, d+1)
	return randomNum
}

func threeDSix() int {
	n := rollDice(6) + rollDice(6) + rollDice(6)
	return n
}

func twoDSix() int {
	n := rollDice(6) + rollDice(6) + 3
	return n
}

func genAbilitiesHuman() abilities {
	a := abilities{threeDSix(), threeDSix(), threeDSix(), threeDSix(), threeDSix(), threeDSix()}
	return a
}

func genAbilitiesDwarf() abilities {
	a := abilities{13, twoDSix(), 13, twoDSix(), twoDSix(), twoDSix()}
	return a
}

func genAbilitiesElf() abilities {
	a := abilities{twoDSix(), 13, twoDSix(), 13, twoDSix(), twoDSix()}
	return a
}

func genAbilitiesHalf() abilities {
	a := abilities{twoDSix(), twoDSix(), twoDSix(), twoDSix(), 13, 13}
	return a
}

func findMod(i int) float64 {
	id := float64(i) / 2
	idf := math.Floor(id)
	idfs := idf - 5
	return idfs
}

func printAbilities(name string, ab int, mod float64) {
	fmt.Println(name + ":")
	fmt.Println(ab)
	fmt.Println(name + " MOD:")
	fmt.Println(mod)
	fmt.Println("")
}

func printCharaHuman(a abilities) {
	fmt.Println("over all quality")
	aSum := a.strength + a.dexterity + a.constitution + a.intelligence + a.wisdom + a.charisma
	fmt.Println(aSum)
	fmt.Println(aSum - 63)
	fmt.Println("")

	printAbilities("STR", a.strength, findMod(a.strength))
	printAbilities("DEX", a.dexterity, findMod(a.dexterity))
	printAbilities("CON", a.constitution, findMod(a.constitution))
	printAbilities("INT", a.intelligence, findMod(a.intelligence))
	printAbilities("WIS", a.wisdom, findMod(a.wisdom))
	printAbilities("CHA", a.charisma, findMod(a.charisma))
}

func printCharaNotHuman(a abilities) {
	fmt.Println("over all quality")
	aSum := a.strength + a.dexterity + a.constitution + a.intelligence + a.wisdom + a.charisma
	fmt.Println(aSum)
	fmt.Println(aSum - 66)
	fmt.Println("")

	printAbilities("STR", a.strength, findMod(a.strength))
	printAbilities("DEX", a.dexterity, findMod(a.dexterity))
	printAbilities("CON", a.constitution, findMod(a.constitution))
	printAbilities("INT", a.intelligence, findMod(a.intelligence))
	printAbilities("WIS", a.wisdom, findMod(a.wisdom))
	printAbilities("CHA", a.charisma, findMod(a.charisma))
}

func main() {

	fmt.Println("--- Hello! ---")
	fmt.Println("I want to make a D&D character for you!")
	fmt.Println("What race would you like to play?")
	fmt.Println("- Human")
	fmt.Println("- Dwarf")
	fmt.Println("- Elf")
	fmt.Println("- Halfling")

	fmt.Print("Enter text: ")
	var input string
	e, _ := fmt.Scanln(&input)
	if input == "Human" {
		fmt.Println("### Here is your new Human Character! ###")

		printCharaHuman(genAbilitiesHuman())
	} else if input == "Dwarf" {
		fmt.Println("### Here is your new Dwarf Character! ###")

		printCharaNotHuman(genAbilitiesDwarf())
	} else if input == "Elf" {
		fmt.Println("### Here is your new Elf Character! ###")

		printCharaNotHuman(genAbilitiesElf())
	} else if input == "Halfing" {
		fmt.Println("### Here is your new Halfing Character! ###")

		printCharaNotHuman(genAbilitiesHalf())
	} else {
		fmt.Println("no!")
		fmt.Println(e)
	}
}
