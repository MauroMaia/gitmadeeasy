package branch

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/MauroMaia/gitmadeeasy/pkg/model"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/jroimartin/gocui"
	"log"
)

var branches []model.Branch

func init() {
	branches = gitcmd.ListBranches()
}

func LayoutListBranches(g *gocui.Gui, xBegins int, yBegins int) *gocui.View {

	var stringLen = 0
	for _, branch := range branches {
		if len(branch.GetName()) > stringLen {
			stringLen = len(branch.GetName())
		}
	}

	_, maxY := g.Size()

	v, err := g.SetView(constants.BRANCH_LIST_VIEW, xBegins, yBegins, xBegins+stringLen+2, maxY-3)
	if err != nil && err != gocui.ErrUnknownView {
		log.Fatalln(err)
	}

	v.Clear()

	for _, value := range branches {
		var name = value.GetName()
		_, _ = fmt.Fprintln(v, name)
	}

	v.Title = "Branch list"

	return v
}

func MenuCursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy+2 > len(branches) {
			// reatch the bottom of the list
			return nil
		}

		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func MenuCursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func Keybindings(g *gocui.Gui) error {

	if err := g.SetKeybinding(constants.BRANCH_LIST_VIEW, gocui.KeyArrowDown, gocui.ModNone, MenuCursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.BRANCH_LIST_VIEW, gocui.KeyArrowUp, gocui.ModNone, MenuCursorUp); err != nil {
		return err
	}
	return nil
}
