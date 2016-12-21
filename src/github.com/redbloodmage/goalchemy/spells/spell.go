package spells

import (
	"github.com/redbloodmage/goalchemy/recipes"
)

type Spell struct {
	Name   string
	Color  string
	Recipe recipes.Recipe
}

func (s *Spell) GetName() (name string) {
	return s.Name
}

func (s *Spell) GetColor() (color string) {
	return s.Color
}

func (s *Spell) GetRecipe() (recipe recipes.Recipe) {
	return s.Recipe
}

func (s *Spell) SetName(name string) {
	s.Name = name
}

func (s *Spell) SetColor(color string) {
	s.Color = color
}

func (s *Spell) SetRecipe(recipe recipes.Recipe) {
	s.Recipe = recipe
}

func NewSpell(name string, color string, recipe recipes.Recipe) *Spell {
	s := &Spell{}
	s.SetName(name)
	s.SetColor(color)
	s.SetRecipe(recipe)
	return s
}
