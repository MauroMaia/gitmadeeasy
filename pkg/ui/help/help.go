package help

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

const HELP_TEXT_MENU = "KeyArrow Down/UP Slect menu option | Enter select Option"
const HELP_TEXT_DEFAULT = "This view does not yet have a help page. More at: https://github.com/MauroMaia/gitmadeeasy"

func LayoutShowHelpView(g *gocui.Gui, xBegins int, yBegins int) *gocui.View {

	maxX, maxY := g.Size()

	v, err := g.SetView(constants.HELP_VIEW, xBegins, yBegins, maxX, maxY)
	if err != nil && err != gocui.ErrUnknownView {
		utils.Logger.Fatalln(err.Error())
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
