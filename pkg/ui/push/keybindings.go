package push

import (
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/jroimartin/gocui"
)

func Keybindings(g *gocui.Gui) error {

	// pop ui to create new branch
	if err := g.SetKeybinding(constants.PUSH_POPUP, gocui.KeyCtrlSpace, gocui.ModNone, quitPopup); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.PUSH_POPUP, gocui.KeyEnter, gocui.ModNone, onEnterPress); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.PUSH_POPUP, gocui.KeySpace, gocui.ModNone, onEnterSpace); err != nil {
		return err
	}

	return nil
}
