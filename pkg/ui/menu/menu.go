package ui

import (
	"fmt"
	"log"

	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

const TOP_MENU = "top_menu"
var buttons = [3]string{
	"New Branch",
	"Committs",
	"List Branchs",
}

func LayoutTopMenuOptions(g *gocui.Gui, xBegins int, yBegins int,yEnd int) *gocui.View {
	var stringLen = 0
	for _, str := range buttons {
		if len(str) > stringLen {
			stringLen = len(str)
		}
	}

	v, err := g.SetView(TOP_MENU, xBegins, yBegins, stringLen+1, yEnd)
	if err != nil && err != gocui.ErrUnknownView {
		log.Fatalln(err)
	}

	v.Clear()

	for _, value := range buttons {
		_, _ = fmt.Fprintln(v, value)
	}

	v.Title="Cmds"

	if _, err = utils.SetCurrentViewOnTop(g, TOP_MENU); err != nil {
		log.Fatalln(err)
	}
	
	return v
}
