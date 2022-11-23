package commit

import (
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

func Keybindings(g *gocui.Gui) error {

	if err := g.SetKeybinding(constants.COMMIT_LIST_VIEW, gocui.KeyArrowDown, gocui.ModNone, menuCursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.COMMIT_LIST_VIEW, gocui.KeyArrowUp, gocui.ModNone, menuCursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.COMMIT_LIST_VIEW, gocui.KeyCtrlSpace, gocui.ModNone, utils.QuitToMenu); err != nil {
		return err
	}

	if err := g.SetKeybinding(constants.DIFF_VIEW, gocui.KeyArrowDown, gocui.ModNone, diffCursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.DIFF_VIEW, gocui.KeyArrowUp, gocui.ModNone, diffCursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.DIFF_VIEW, gocui.KeyCtrlSpace, gocui.ModNone, utils.QuitToMenu); err != nil {
		return err
	}

	return nil
}
