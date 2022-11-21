package branch

import (
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/jroimartin/gocui"
)

func Keybindings(g *gocui.Gui) error {

	if err := g.SetKeybinding(constants.BRANCH_LIST_VIEW, gocui.KeyArrowDown, gocui.ModNone, listCursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.BRANCH_LIST_VIEW, gocui.KeyArrowUp, gocui.ModNone, listCursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.NEW_COMMIT_POPUP, gocui.KeyEnter, gocui.ModNone, onEnterPress); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.NEW_COMMIT_POPUP, gocui.KeyCtrlSpace, gocui.ModNone, Popoff); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.NEW_COMMIT_POPUP, gocui.KeySpace, gocui.ModNone, ignoreKey); err != nil {
		return err
	}
	return nil
}
