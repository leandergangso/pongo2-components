package pongo2components

import (
	"fmt"

	"github.com/flosch/pongo2/v6"
)

type componentTag struct {
	comp Component
	// wrappers []*pongo2.NodeWrapper
	// slots    []Slot
}

func (tag componentTag) Execute(ctx *pongo2.ExecutionContext, w pongo2.TemplateWriter) *pongo2.Error {
	f, err := pongo2.DefaultSet.FromFile(componentDirPath + "/" + tag.comp.Name + ".html")
	if err != nil {
		return ctx.OrigError(err, nil)
	}

	ctx.Private[componentCtxKey] = tag.comp
	// todo set rest of ctx

	err = f.ExecuteWriter(ctx.Private, w)
	if err != nil {
		return ctx.OrigError(err, nil)
	}

	return nil
}

func componentTagParser(doc *pongo2.Parser, start *pongo2.Token, args *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	tag := new(componentTag)

	tok := args.Current()

	if tok == nil || tok.Val == "" {
		return tag, &pongo2.Error{
			OrigError: fmt.Errorf("expected component name as first argument, got (empty string)"),
		}
	}

	comp, ok := componentsRegister[tok.Val]
	if !ok {
		return tag, &pongo2.Error{
			OrigError: fmt.Errorf("component '%s' is not registered", tok.Val),
		}
	}
	tag.comp = comp

	// dtok := doc.Current()
	// fmt.Println("start", dtok.Val, "end")
	// if dtok.Val == "" {
	// 	panic("no slot data")
	// }

	// args.Consume()

	insideSlot := false

	for {
		wrapper, tagArgs, err := doc.WrapUntilTag(slotTagName, endSlotTagName, endComponentTagName)
		if err != nil {
			return tag, err
		}
		// tag.wrappers = append(tag.wrappers, wrapper)

		fmt.Println(tagArgs.Current())

		if wrapper.Endtag == slotTagName {
			if insideSlot {
				return tag, tagArgs.Error("slot tag not allowed here", nil)
			}
			insideSlot = true

			// panic("here")

			// tagArgs.
		} else {
			if tagArgs.Count() > 0 {
				// end tags for slot and component can't take any arguments
				return tag, tagArgs.Error("arguments not allowed here", nil)
			}
			if wrapper.Endtag == endSlotTagName {
				insideSlot = false
			}
		}

		if wrapper.Endtag == endComponentTagName {
			break
		}
	}

	// props := make([]string, len(comp.Props))
	// slots := make([]string, len(comp.Slots))

	// for args.Remaining() != 0 {
	// 	tok = args.Current()

	// 	fmt.Println("here", tok)

	// 	args.Consume()
	// }

	// check required props/slots
	// todo
	// for _, prop := range comp.Props {
	// }

	// check undefined props/slots
	// todo

	return tag, nil
}
