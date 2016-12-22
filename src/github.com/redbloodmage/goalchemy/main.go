package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/redbloodmage/goalchemy/components"
	"github.com/redbloodmage/goalchemy/recipes"
	"github.com/redbloodmage/goalchemy/spells"
)

var (
	compSlice = make([]components.Component, 5)
	spellList = make([]spells.Spell, 2)
	grimore   = make([]recipes.Recipe, 2)
)

func primaryRecursion(prompt string) {

	fmt.Println("Started primaryRecursion()")
	//	var mPromptOne string
	//	var mPromptTwo string

	switch prompt {

	case "q", "Q", "quit", "QUIT":
		os.Exit(0)
	case "C", "c":
		fmt.Println("C Selected:")
		fmt.Println(compSlice)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Choose one of the following: ")
		fmt.Println("(C)omponents,(S)pellist,(G)rimore, (M)erge (Q)uit")
		scanner.Scan()
		prompt = scanner.Text()
		primaryRecursion(prompt)
	case "S", "s":
		fmt.Println("S selected")
		fmt.Println(spellList)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Choose one of the following: ")
		fmt.Println("(C)omponents,(S)pellist,(G)rimore, (M)erge (Q)uit")
		scanner.Scan()
		prompt = scanner.Text()
		primaryRecursion(prompt)
	case "G", "g":
		fmt.Println("G Selected:")
		fmt.Println(grimore)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Choose one of the following: ")
		fmt.Println("(C)omponents,(S)pellist,(G)rimore, (M)erge (Q)uit")
		scanner.Scan()
		prompt = scanner.Text()
		primaryRecursion(prompt)

	case "M", "m":
		fmt.Println("M selected")
		/*
			Select the first component from the compSlice.  Save it to a var.  Select the second component
			from the compSlice and save that to another var.  When both are selected send it to a combine method
		*/
		Mscanner := bufio.NewScanner(os.Stdin)

		//TODO: START HERE NEXT REVISIT
		fmt.Println("Select the First Component to Merge.", compSlice)
		Mscanner.Scan()
		//		mPromptOne = Mscanner.Text()
		arr := []string{"test", "some_other_test", "foo_etc"}
		nofoos := []string{}
		for i := range arr {
			if !strings.HasPrefix(arr[i], "foo_") && len(arr) <= 7 {
				nofoos = append(nofoos, arr[i])
			}
		}
		fmt.Println(nofoos) // [test]
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Choose one of the following: ")
		fmt.Println("(C)omponents,(S)pellist,(G)rimore, (M)erge (Q)uit")
		scanner.Scan()
		prompt = scanner.Text()
		primaryRecursion(prompt)

	default:
		fmt.Println("Default Fired")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Choose one of the following: ")
		fmt.Println("(C)omponents,(S)pellist,(G)rimore, (M)erge (Q)uit")
		scanner.Scan()
		prompt = scanner.Text()
		primaryRecursion(prompt)

	}
} //End primaryRecursion

//TODO: Create mergeComponents function
//func mergeComponents()

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

	//Some appended items are pointers vs. literal depending on how I made them. Setters vs. Constructor
	compSlice = append(compSlice, SpringWater, Flame, *Sand, *SunBeam, *Snow)

	//Creating the RegenLife Recipe
	regenLifeComps := make([]components.Component, 5)
	regenLifeComps = append(regenLifeComps, SpringWater, *SunBeam)
	RegenLifeRecipe := recipes.NewRecipe("Recipe of Regeneration", regenLifeComps, "Green")

	//Creating RegenLife Spell
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

	//Creating Spell list aggregating all spells into one slice
	spellList = append(spellList, FlameBeam, RegenLife)
	//Creating Grimore (Recipe List) aggregating all recipes
	grimore = append(grimore, *RegenLifeRecipe, *FlameBeamRecipe)
	//compMap := make(map[string][]Component)

	var initialPrompt string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Choose one of the following: ")
	fmt.Println("(C)omponents,(S)pellist,(G)rimore, (M)erge (Q)uit")
	scanner.Scan()
	initialPrompt = scanner.Text()
	primaryRecursion(initialPrompt)

} // END MAIN
