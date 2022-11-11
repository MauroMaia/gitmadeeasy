package help

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/jroimartin/gocui"
	"log"
)

const HELP_TEXT_MENU = "KeyArrow Down/UP Slect menu option | Enter select Option"
const HELP_TEXT_DEFAULT = "This view does not yet have a help page. More at: https://github.com/MauroMaia/gitmadeeasy"

func LayoutShowHelpView(g *gocui.Gui, xBegins int, yBegins int) *gocui.View {

	//log.Fatalf("xBegins %d yBegins%d \n",xBegins,yBegins)

	maxX, maxY := g.Size()
	//log.Fatalf("maxX %d stringLen %d xBegins %d\n",maxX,stringLen,xBegins)

	v, err := g.SetView(constants.HELP_VIEW, xBegins, yBegins, maxX, maxY)
	if err != nil && err != gocui.ErrUnknownView {
		log.Fatalln(err)
	}

	v.Clear()

	switch constants.SELECTED_MENU {
	case constants.MENU_VIEW:
		_, _ = fmt.Fprintln(v, HELP_TEXT_MENU)
		break
	default:
		_, _ = fmt.Fprintln(v, HELP_TEXT_DEFAULT)
	}

	v.Title = "Help"

	return v
}
