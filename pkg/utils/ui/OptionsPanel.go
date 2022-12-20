package ui

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

type OptionsElementType interface {
	bool | int
}

type OptionsElement[T OptionsElementType] struct {
	key         string
	description string
	Val         T
}

func NewOptionsElement[T OptionsElementType](key string, description string, val T) OptionsElement[T] {
	return OptionsElement[T]{key, description, val}
}

//
//	OptionsPanel
//

type OptionsPanel struct {
	name    string
	tile    string
	x, y    int
	h, w    int
	options *[]interface{}
}

func NewOptionsPanel(name string, title string, x int, y int, w int, options *[]interface{}) OptionsPanel {
	return OptionsPanel{name, title, x, y, len(*options)*2 + 1, w, options}
}

func (p OptionsPanel) Layout(g *gocui.Gui) error {

	if v, err := g.SetView(p.name, p.x, p.y, p.x+p.w, p.y+p.h); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Clear()

		for _, option := range *p.options {
			switch option.(type) {
			case OptionsElement[bool]:
				optionConv := option.(OptionsElement[bool])
				if optionConv.Val == true {
					_, err = fmt.Fprintf(v, "[X] %s\n", optionConv.key)
				} else {
					_, err = fmt.Fprintf(v, "[ ] %s\n", optionConv.key)
				}
				_, err = fmt.Fprintf(v, "  -- %s\n", optionConv.description)

			default:
				utils.Logger.Fatalf("unknown variable type")
			}
		}

		v.Editable = false
		v.Wrap = false // may cause unexpected errors on click up/down
		v.Title = p.tile
	}
	return nil
}

func (p OptionsPanel) Update(g *gocui.Gui, options *[]interface{}) error {
	p.options = options

	var err error

	if v, err := g.View(p.name); v != nil {
		if err == gocui.ErrUnknownView {
			return err
		}

		v.Clear()

		for _, option := range *p.options {
			switch option.(type) {
			case OptionsElement[bool]:
				optionConv := option.(OptionsElement[bool])
				if optionConv.Val == true {
					_, err = fmt.Fprintf(v, "[X] %s\n", optionConv.key)
				} else {
					_, err = fmt.Fprintf(v, "[ ] %s\n", optionConv.key)
				}
				_, err = fmt.Fprintf(v, "  -- %s\n", optionConv.description)

			default:
				utils.Logger.Fatalf("unknown variable type")
			}
		}
	}
	return err
}
