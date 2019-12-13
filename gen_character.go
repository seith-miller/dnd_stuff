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

func main() {
	fmt.Println("### Here is your new Human Character! ###")

	a := genAbilitiesHuman()

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

	fmt.Println("### Here is your new Dwarf Character! ###")

	b := genAbilitiesDwarf()

	fmt.Println("over all quality")
	bSum := b.strength + b.dexterity + b.constitution + b.intelligence + b.wisdom + b.charisma
	fmt.Println(bSum)
	fmt.Println(bSum - 66)
	fmt.Println("")

	printAbilities("STR", b.strength, findMod(b.strength))
	printAbilities("DEX", b.dexterity, findMod(b.dexterity))
	printAbilities("CON", b.constitution, findMod(b.constitution))
	printAbilities("INT", b.intelligence, findMod(b.intelligence))
	printAbilities("WIS", b.wisdom, findMod(b.wisdom))
	printAbilities("CHA", b.charisma, findMod(b.charisma))
}
