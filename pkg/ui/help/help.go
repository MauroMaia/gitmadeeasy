package help

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

const HELP_TEXT_MENU = "KeyArrow Down/UP Slect menu option | Enter select Option"

const HELP_VIEW = "HELP_VIEW"

func LayoutShowHelpView(g *gocui.Gui, xBegins int, yBegins int) *gocui.View {

	//log.Fatalf("xBegins %d yBegins%d \n",xBegins,yBegins)

	maxX, maxY := g.Size()
	//log.Fatalf("maxX %d stringLen %d xBegins %d\n",maxX,stringLen,xBegins)

	v, err := g.SetView(HELP_VIEW, xBegins, yBegins, maxX-1, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		log.Fatalln(err)
	}

	v.Clear()

	_, _ = fmt.Fprintln(v, HELP_TEXT_MENU)

	v.Title = "Help"

	return v
}
