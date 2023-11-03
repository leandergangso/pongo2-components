package main

import (
	"fmt"

	pongo2components "github.com/leandergangso/pongo2-components"
)

func init() {
	fmt.Println("main")
	pongo2components.Register(pongo2components.Component{
		Name:     "button",
		FilePath: "button.html",
		Props:    []pongo2components.Prop{"value", "type"},
		Slots: []pongo2components.Slot{
			{
				Name: pongo2components.SlotDefault,
			},
		},
	})
}

func main() {
	pongo2components.Init()
}
