package pull

import (
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/jroimartin/gocui"
)

func Keybindings(g *gocui.Gui) error {

	// pop ui to create new branch
	if err := g.SetKeybinding(constants.PUll_POPUP, gocui.KeyCtrlSpace, gocui.ModNone, quitPopup); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.PUll_POPUP, gocui.KeyEnter, gocui.ModNone, onEnterPress); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.PUll_POPUP, gocui.KeySpace, gocui.ModNone, onEnterSpace); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.PUll_POPUP, gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.PUll_POPUP, gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}

	return nil
}
