package commit

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

func DisplayPopUp(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(constants.COMMIT_POPUP, maxX/3-30, maxY/3, 2*(maxX/3)+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		if _, err := g.SetCurrentView(constants.COMMIT_POPUP); err != nil {
			return err
		}
		v.Editable = true
		v.Title = "Commit Message"

		if _, err := utils.SetCurrentViewOnTop(g, constants.COMMIT_POPUP); err != nil {
			utils.Logger.Fatalln(err)
		}
		constants.HIGHLIGHT_VIEW = constants.COMMIT_POPUP
	}
	return nil
}

func OpenPopup(g *gocui.Gui, v *gocui.View) error {
	DisplayPopUp(g)
	constants.HIGHLIGHT_VIEW = constants.COMMIT_POPUP
	return nil
}

func quitPopup(g *gocui.Gui, v *gocui.View) error {

	content := v.ViewBuffer()
	utils.Logger.WithField("commitMessage", content)
	_, err := gitcmd.Commit(content, false)
	if err != nil {
		v.Clear()
		v.Editable = false
		v.BgColor = gocui.ColorRed
		v.Rewind()
		_, _ = fmt.Fprintln(v, err.Error())
		return err
	}

	go LoadCommits()

	g.DeleteView(constants.COMMIT_POPUP)
	constants.LEFT_VIEW = constants.COMMIT_LIST_VIEW
	constants.HIGHLIGHT_VIEW = constants.COMMIT_LIST_VIEW

	return nil
}
