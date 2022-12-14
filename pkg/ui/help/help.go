package help

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

const HELP_TEXT_MENU = "KeyArrow Down/UP Slect menu option | Enter select Option"
const HELP_TEXT_LIST = "KeyArrow Down/UP to scroll | Press Crontol + space to go back to menu"
const HELP_TEXT_FILE_CHANGED = "KeyArrow Down/UP to scroll | Press Crontol + S to add file | Press Crontol + X write commit"
const HELP_TEXT_DEFAULT = "This view(%s) does not yet have a help information. More at: https://github.com/MauroMaia/gitmadeeasy"

func LayoutShowHelpView(g *gocui.Gui, xBegins int, yBegins int) *gocui.View {

	maxX, maxY := g.Size()

	v, err := g.SetView(constants.HELP_VIEW, xBegins, yBegins, maxX, maxY)
	if err != nil && err != gocui.ErrUnknownView {
		utils.Logger.Fatalln(err.Error())
	}

	v.Clear()

	switch constants.HIGHLIGHT_VIEW {
	case constants.MENU_VIEW:
		_, _ = fmt.Fprintln(v, HELP_TEXT_MENU)
		break
	case constants.BRANCH_LIST_VIEW:
		_, _ = fmt.Fprintln(v, HELP_TEXT_LIST)
		break
	case constants.COMMIT_LIST_VIEW:
		_, _ = fmt.Fprintln(v, HELP_TEXT_LIST)
		break
	case constants.FILE_CHANGED_VIEW:
		_, _ = fmt.Fprintln(v, HELP_TEXT_FILE_CHANGED)
		break

	default:
		_, _ = fmt.Fprintf(v, HELP_TEXT_DEFAULT, constants.HIGHLIGHT_VIEW)
	}

	v.Title = "Help"

	return v
}
