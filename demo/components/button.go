package components

import (
	"github.com/leandergangso/pongo2components"
)

func init() {
	pongo2components.Register(pongo2components.Component{
		Name: "button",
		Props: []pongo2components.Prop{
			{
				Name:     "value",
				Required: true,
			},
			{
				Name:         "type",
				DefaultValue: "primary",
			},
		},
		Slots: []pongo2components.Slot{
			{
				Name:    "default",
				Default: true,
			},
			{
				Name: "before",
			},
		},
	})
}
