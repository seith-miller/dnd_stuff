package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type abilityies struct {
	strength     int
	dexterity    int
	constitution int
	intelligence int
	wisdom       int
	charisma     int
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

func genAbilitiesHuman() abilityies {
	a := abilityies{threeDSix(), threeDSix(), threeDSix(), threeDSix(), threeDSix(), threeDSix()}
	return a
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func findMod(i int) float64 {
	id := float64(i) / 2
	idf := math.Floor(id)
	idfs := idf - 5
	return idfs
}

func main() {
	fmt.Println("### Here is your new Character! ###")

	a := genAbilitiesHuman()

	fmt.Println("over all quality")
	aSum := a.strength + a.dexterity + a.constitution + a.intelligence + a.wisdom + a.charisma
	fmt.Println(aSum)
	fmt.Println(aSum - 63)
	fmt.Println("")

	fmt.Println("STR:")
	fmt.Println(a.strength)
	fmt.Println("STR mod:")
	fmt.Println(findMod(a.strength))
	fmt.Println("")

	fmt.Println("DEX:")
	fmt.Println(a.dexterity)
	fmt.Println("DEX mod:")
	fmt.Println(findMod(a.dexterity))
	fmt.Println("")

	fmt.Println("CON:")
	fmt.Println(a.constitution)
	fmt.Println("CON mod:")
	fmt.Println(findMod(a.constitution))
	fmt.Println("")

	fmt.Println("INT:")
	fmt.Println(a.intelligence)
	fmt.Println("INT mod:")
	fmt.Println(findMod(a.intelligence))
	fmt.Println("")

	fmt.Println("WIS:")
	fmt.Println(a.wisdom)
	fmt.Println("WIS mod:")
	fmt.Println(findMod(a.wisdom))
	fmt.Println("")

	fmt.Println("CHA:")
	fmt.Println(a.charisma)
	fmt.Println("CHA mod:")
	fmt.Println(findMod(a.charisma))
	fmt.Println("")
}
