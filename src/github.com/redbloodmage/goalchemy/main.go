package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/redbloodmage/goalchemy/components"
	"github.com/redbloodmage/goalchemy/recipes"
	"github.com/redbloodmage/goalchemy/spells"
)

var (
	compSlice = make([]components.Component, 5)
)

func primaryRecursion(prompt string) {

	fmt.Println("Started primaryRecursion()")

	switch prompt {

	case "q", "Q", "quit", "QUIT":
		os.Exit(0)
	case "C", "c":
		fmt.Println("C Selected:")
		fmt.Println(compSlice)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Choose one of the following: ")
		fmt.Println("(C)omponents,(S)pellist,(G)rimore, (Q)uit")
		scanner.Scan()
		prompt = scanner.Text()
		primaryRecursion(prompt)
	case "S", "s":
		fmt.Println("S selected")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Choose one of the following: ")
		fmt.Println("(C)omponents,(S)pellist,(G)rimore, (Q)uit")
		scanner.Scan()
		prompt = scanner.Text()
		primaryRecursion(prompt)
	case "G", "g":
		fmt.Println("G selected")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Choose one of the following: ")
		fmt.Println("(C)omponents,(S)pellist,(G)rimore, (Q)uit")
		scanner.Scan()
		prompt = scanner.Text()
		primaryRecursion(prompt)
	default:
		fmt.Println("Default Fired")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Choose one of the following: ")
		fmt.Println("(C)omponents,(S)pellist,(G)rimore, (Q)uit")
		scanner.Scan()
		prompt = scanner.Text()
		primaryRecursion(prompt)

	}
} //End primaryRecursion

func main() {

	//Creating SpringWater Component with Setters
	SpringWater := components.Component{}
	SpringWater.SetName("SpringWater")
	SpringWater.SetColor("Blue")
	SpringWater.SetElement("Water")

	//Creating Sand Component using the Constructor
	Sand := components.NewComponent("Sand", "Red", "Earth")

	//Creating SunBeam Component by setting public attributes direct
	SunBeam := components.NewComponent("Sunbeam", "Red", "Earth")

	//Creating Flame Component
	Flame := components.Component{
		Name:    "Flame",
		Element: "Fire",
		Color:   "Red",
	}
	//Creating Snow Component via Method
	Snow := components.NewComponent("Snow", "Blue", "Water")

	compSlice = append(compSlice, SpringWater, Flame, *Sand, *SunBeam, *Snow)
	//Creating the RegenLife Recipe
	x := make([]components.Component, 5)
	x = append(x, SpringWater, *SunBeam)
	RegenLifeRecipe := recipes.NewRecipe("Recipe of Regeneration", x, "Green")

	//Creating Regeneration Spell
	RegenLife := spells.Spell{}
	RegenLife.SetName("Recipe of Regeneration")
	RegenLife.SetColor("Green")
	RegenLife.SetRecipe(*RegenLifeRecipe)

	//Creating the Flamebeam Components and Recipe
	FlameBeamComps := make([]components.Component, 5)
	FlameBeamComps = append(FlameBeamComps, *SunBeam, Flame)
	FlameBeamRecipe := recipes.NewRecipe("Recipe of Flamebeam", FlameBeamComps, "Red")
	fmt.Println("Flamebeam recipe:", FlameBeamRecipe)

	//Creating Flamebeam Spell
	FlameBeam := spells.Spell{}
	FlameBeam.SetName("Recipe of Flamebeam")
	FlameBeam.SetColor("Red")
	FlameBeam.SetRecipe(*FlameBeamRecipe)

	// Printing out the creations (Spells, Components, etc.)
	fmt.Println("\nComponents:\n")
	fmt.Println(SpringWater)
	fmt.Println(SunBeam)
	fmt.Println(Flame)

	fmt.Println("\nSpells:\n")
	fmt.Println(RegenLife)
	fmt.Println(FlameBeam)

	//compMap := make(map[string][]Component)

	var initialPrompt string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Choose one of the following: ")
	fmt.Println("(C)omponents,(S)pellist,(G)rimore, (Q)uit")
	scanner.Scan()
	initialPrompt = scanner.Text()
	primaryRecursion(initialPrompt)

} // END MAIN
