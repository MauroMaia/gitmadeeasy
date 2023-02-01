package log

import (
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

func Keybindings(g *gocui.Gui) error {

	//
	//	LOG_VIEW
	//
	if err := g.SetKeybinding(constants.LOG_VIEW, gocui.KeyArrowDown, gocui.ModNone, menuCursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.LOG_VIEW, gocui.KeyArrowUp, gocui.ModNone, menuCursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.LOG_VIEW, gocui.KeyCtrlSpace, gocui.ModNone, utils.QuitToMenu); err != nil {
		return err
	}

	return nil
}
