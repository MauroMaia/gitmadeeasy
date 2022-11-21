package utils

import (
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/jroimartin/gocui"
)

func SetCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func QuitToMenu(g *gocui.Gui, v *gocui.View) error {

	constants.SELECTED_MENU = constants.MENU_VIEW

	return nil
}
