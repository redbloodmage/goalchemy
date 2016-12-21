package components

import ()

type Component struct {
	Name    string
	Element string
	Color   string
}

func (c *Component) SetName(name string) {
	c.Name = name
}

func (c *Component) SetElement(element string) {
	c.Element = element
}

func (c *Component) SetColor(color string) {
	c.Color = color
}

func (c *Component) GetName() string {
	return c.Name
}

func (c *Component) GetColor() string {
	return c.Color
}

func (c *Component) GetElement() string {
	return c.Element
}

func NewComponent(name string, color string, element string) *Component {
	c := &Component{}
	c.SetName(name)
	c.SetColor(color)
	c.SetElement(element)
	return c
}
