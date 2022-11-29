package branch

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

func DisplayPopUp(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(constants.NEW_BRANCH_POPUP, maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		if _, err := g.SetCurrentView(constants.NEW_BRANCH_POPUP); err != nil {
			return err
		}
		v.Editable = true
		v.Title = "Write new branch name"

		if _, err := utils.SetCurrentViewOnTop(g, constants.NEW_BRANCH_POPUP); err != nil {
			utils.Logger.Fatalln(err)
		}
		constants.HIGHLIGHT_VIEW = constants.NEW_BRANCH_POPUP
	}
	return nil
}

func onEnterPress(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}
	if len(l) > 0 {
		err = gitcmd.CreateNewBranch(l, false)
		if err != nil {
			v.Clear()
			v.Editable = false
			v.BgColor = gocui.ColorRed
			v.Rewind()
			_, _ = fmt.Fprintln(v, err.Error())
			v.SetCursor(0, 46)
		} else {
			RefreshBranchList()
			g.DeleteView(constants.NEW_BRANCH_POPUP)
			constants.LEFT_VIEW = constants.BRANCH_LIST_VIEW
			constants.HIGHLIGHT_VIEW = constants.MENU_VIEW
		}

	}
	return nil
}

func quitPopup(g *gocui.Gui, v *gocui.View) error {

	g.DeleteView(constants.NEW_BRANCH_POPUP)
	constants.HIGHLIGHT_VIEW = constants.MENU_VIEW

	return nil
}

func ignoreKey(g *gocui.Gui, v *gocui.View) error {
	return nil
}
