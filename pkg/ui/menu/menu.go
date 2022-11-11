package ui

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"log"

	"github.com/jroimartin/gocui"
)

const B_NEW_BRANCH = "New Branch"
const B_SHOW_COMMITS = "Commit's List"
const B_SHOW_BRANCHS = "Branch's List"

var buttons = [3]string{
	B_NEW_BRANCH,
	B_SHOW_COMMITS,
	B_SHOW_BRANCHS,
}

func LayoutTopMenuOptions(g *gocui.Gui, xBegins int, yBegins int, yEnd int) *gocui.View {
	var stringLen = 0
	for _, str := range buttons {
		if len(str) > stringLen {
			stringLen = len(str)
		}
	}

	v, err := g.SetView(constants.MENU_VIEW, xBegins, yBegins, stringLen+1, yEnd)
	if err != nil && err != gocui.ErrUnknownView {
		log.Fatalln(err)
	}

	v.Clear()

	for _, value := range buttons {
		_, _ = fmt.Fprintln(v, value)
	}

	v.Title = "MENU"
	v.Highlight = true

	return v
}

func MenuCursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy+2 > len(buttons) {
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

func getLine(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	switch l {
	case B_NEW_BRANCH:
		log.Fatal(l)
		break
	case B_SHOW_COMMITS:
		if _, err = utils.SetCurrentViewOnTop(g, constants.COMMIT_LIST_VIEW); err != nil {
			log.Fatalln(err)
		}
		constants.SELECTED_MENU = constants.COMMIT_LIST_VIEW
		break
	case B_SHOW_BRANCHS:
		if _, err = utils.SetCurrentViewOnTop(g, constants.BRANCH_LIST_VIEW); err != nil {
			log.Fatalln(err)
		}
		constants.SELECTED_MENU = constants.BRANCH_LIST_VIEW
		break
	default:
		// not expected do nothing
		return nil
	}

	return nil
}

func Keybindings(g *gocui.Gui) error {

	if err := g.SetKeybinding(constants.MENU_VIEW, gocui.KeyArrowDown, gocui.ModNone, MenuCursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.MENU_VIEW, gocui.KeyArrowUp, gocui.ModNone, MenuCursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.MENU_VIEW, gocui.KeyEnter, gocui.ModNone, getLine); err != nil {
		return err
	}
	return nil
}
