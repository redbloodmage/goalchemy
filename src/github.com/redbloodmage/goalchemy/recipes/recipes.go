package recipes

import (
	"github.com/redbloodmage/goalchemy/components"
)

type Recipe struct {
	Name        string
	Ingredients *[]components.Component
	Color       string
}

func (r *Recipe) GetName() (name string) {
	return r.Name
}

func (r *Recipe) GetIngredients() (ingredients *[]components.Component) {
	return r.Ingredients
}

func (r *Recipe) GetColor() (color string) {
	return r.Color
}

func (r *Recipe) SetName(name string) {
	r.Name = name
}

func (r *Recipe) SetColor(color string) {
	r.Color = color
}

func (r *Recipe) SetIngredients(ingredients []components.Component) {
	r.Ingredients = &ingredients
}

func NewRecipe(name string, ingredients []components.Component, color string) *Recipe {
	r := &Recipe{}
	r.SetName(name)
	r.SetIngredients(ingredients)
	r.SetColor(color)
	return r
}
