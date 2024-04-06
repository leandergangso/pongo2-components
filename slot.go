package pongo2components

import (
	"fmt"

	"github.com/flosch/pongo2/v6"
)

type slotTag struct {
	name  string
	value string
}

func (tag slotTag) Execute(ctx *pongo2.ExecutionContext, w pongo2.TemplateWriter) *pongo2.Error {
	// todo write() default value if no content is defined

	// val, ok := ctx.Private[componentCtxKey]
	// if !ok {
	// 	// todo throw error?
	// 	return nil
	// }

	// comp, ok := val.(Component)
	// if !ok {
	// 	return ctx.Error(fmt.Sprintf("invalid context type for key '%s' of type '%T'", componentCtxKey, val), nil)
	// }

	// for _, slot := range comp.Slots {
	// 	if tag.name == slot.Name {
	// 		w.WriteString(tag.value)
	// 		break
	// 	}
	// }

	return nil
}

func slotTagParser(doc *pongo2.Parser, start *pongo2.Token, args *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	tag := new(slotTag)

	tok := args.Current()

	if tok == nil || tok.Val == "" {
		return tag, &pongo2.Error{
			OrigError: fmt.Errorf("expected slot name as first argument, got (empty string)"),
		}
	}
	tag.name = tok.Val

	_, tagArgs, err := doc.WrapUntilTag(endSlotTagName)
	if err != nil {
		return tag, err
	}

	fmt.Println("slot args", tagArgs.Current())

	return tag, nil
}
