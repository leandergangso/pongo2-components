package pongo2components

import (
	"errors"
	"fmt"
	"strings"

	"github.com/flosch/pongo2/v6"
)

var componentsRegister map[string]Component

type ComponentsTagOptions struct {
	// Name changes the default tag name to the given name.
	//	Default: "component"
	Name string

	// Force will replace the tag with the same name if the tag name already exists.
	//	Default: false
	Force bool
}

// Init will initialize the component tag for use in pongo2 templating.
func Init() error {
	return InitWithOptions(ComponentsTagOptions{
		Name:  "component",
		Force: false,
	})
}

// InitWithOptions will initialize the component tag for use in pongo2 templating.
func InitWithOptions(opt ComponentsTagOptions) error {
	err := pongo2.RegisterTag(opt.Name, tagParser)
	if opt.Force && err != nil {
		err = pongo2.ReplaceTag(opt.Name, tagParser)
		if err != nil {
			return err
		}
	}

	return err
}

// type pongo2component interface {
// 	component()
// }

type Component struct {
	Name     string
	FilePath string
	Props    []Prop
	Slots    []Slot
}

type Prop string

type Slot struct {
	Name string
}

func Register(c Component) error {
	if c.Name == "" {
		return errors.New("name cannot be empty")
	}
	if c.FilePath == "" {
		return errors.New("filePath cannot be empty")
	}

	c.Name = strings.ToLower(c.Name)

	_, ok := componentsRegister[c.Name]
	if !ok {
		return fmt.Errorf("component with same name '%s' is already registered", c.Name)
	}

	return nil
}
