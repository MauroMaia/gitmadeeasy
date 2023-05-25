package commit

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

var commitsIds []string
var posListView = 0

func init() {
	LoadCommits()
}

// TODO - fill the docs
func LoadCommits() {
	commitsIds, _ = gitcmd.ListCommits()
}

func LayoutListCommits(g *gocui.Gui, xBegins int, yBegins int, xEnd int) *gocui.View {

	_, maxY := g.Size()

	v, err := g.SetView(constants.COMMIT_LIST_VIEW, xBegins, yBegins, xEnd, maxY-2)
	if err != nil && err != gocui.ErrUnknownView {
		utils.Logger.Fatalln(err)
	}

	v.Clear()

	for _, value := range commitsIds {
		_, _ = fmt.Fprintln(v, value)
	}

	v.Title = "Last Commits"
	// TODO - create an option in Settings
	v.Wrap = true

	return v
}

func menuCursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if posListView+2 > len(commitsIds) {
			// reach the bottom of the list
			return nil
		}

		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
		posListView++
	}
	return nil
}

func menuCursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
		posListView--
	}
	return nil
}
