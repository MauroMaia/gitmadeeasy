package commit

import (
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/jroimartin/gocui"
)

func Keybindings(g *gocui.Gui) error {

	if err := g.SetKeybinding(constants.COMMIT_LIST_VIEW, gocui.KeyArrowDown, gocui.ModNone, MenuCursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.COMMIT_LIST_VIEW, gocui.KeyArrowUp, gocui.ModNone, MenuCursorUp); err != nil {
		return err
	}
	return nil
}
