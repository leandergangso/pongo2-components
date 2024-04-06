package pongo2components

import (
	"errors"
	"fmt"
	"strings"

	"github.com/flosch/pongo2/v6"
)

var (
	componentsRegister = make(map[string]Component)
	componentDirPath   = ""
)

func Init(dirPath string) {
	// component tag
	err := pongo2.RegisterTag(componentTagName, componentTagParser)
	if err != nil {
		err = pongo2.ReplaceTag(componentTagName, componentTagParser)
		if err != nil {
			panic(err)
		}
	}

	// slot tag
	err = pongo2.RegisterTag(slotTagName, slotTagParser)
	if err != nil {
		err = pongo2.ReplaceTag(slotTagName, slotTagParser)
		if err != nil {
			panic(err)
		}
	}

	componentDirPath = dirPath
}

type Component struct {
	Name  string
	Props []Prop
	Slots []Slot
}

type Prop struct {
	Name         string
	DefaultValue string
	Required     bool
}

type Slot struct {
	Name string

	// Only one slot can have default as true.
	Default bool
}

func Register(c Component) {
	if c.Name == "" {
		panic(errors.New("name cannot be empty"))
	}

	c.Name = strings.ToLower(c.Name)

	_, ok := componentsRegister[c.Name]
	if ok {
		panic(fmt.Errorf("component with same name '%s' is already registered", c.Name))
	}

	componentsRegister[c.Name] = c
}
