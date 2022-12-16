package push

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

var Force = false

func draw(v *gocui.View) error {

	var err error

	v.Clear()

	if Force {
		_, err = fmt.Fprintln(v, "[X] Force")
	} else {
		_, err = fmt.Fprintln(v, "[ ] Force")
	}

	return err
}

func DisplayPopUp(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(constants.PUSH_POPUP, maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		_ = draw(v)

		v.Editable = false
		v.Title = "Push"

		if _, err := g.SetCurrentView(constants.PUSH_POPUP); err != nil {
			return err
		}

		if _, err := utils.SetCurrentViewOnTop(g, constants.PUSH_POPUP); err != nil {
			utils.Logger.Fatalln(err)
		}

		constants.HIGHLIGHT_VIEW = constants.PUSH_POPUP
	}
	return nil
}

func onEnterPress(g *gocui.Gui, v *gocui.View) error {

	_, err := gitcmd.Push(Force)
	if err != nil {
		v.Clear()
		v.Editable = false
		v.BgColor = gocui.ColorRed
		v.Rewind()
		_, _ = fmt.Fprintln(v, err.Error())
		v.SetCursor(0, 0)
	} else {
		g.DeleteView(constants.PUSH_POPUP)
		constants.LEFT_VIEW = constants.COMMIT_LIST_VIEW
		constants.RIGTH_VIEW = constants.BRANCH_LIST_VIEW
		constants.HIGHLIGHT_VIEW = constants.MENU_VIEW
	}

	return nil
}

func onEnterSpace(g *gocui.Gui, v *gocui.View) error {

	Force = !Force

	return draw(v)
}

func quitPopup(g *gocui.Gui, v *gocui.View) error {

	g.DeleteView(constants.PUSH_POPUP)
	constants.HIGHLIGHT_VIEW = constants.MENU_VIEW

	return nil
}
