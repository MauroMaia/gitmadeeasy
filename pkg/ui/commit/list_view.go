package commit

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

var commitsIds []string
var pos = 0

func init() {
	commitsIds = gitcmd.ListCommitIDs()
}

func LayoutListCommits(g *gocui.Gui, xBegins int, yBegins int) *gocui.View {

	var stringLen = len(commitsIds[0])

	_, maxY := g.Size()

	v, err := g.SetView(constants.COMMIT_LIST_VIEW, xBegins, yBegins, xBegins+stringLen+2, maxY-3)
	if err != nil && err != gocui.ErrUnknownView {
		utils.Logger.Fatalln(err)
	}

	v.Clear()

	for _, value := range commitsIds {
		_, _ = fmt.Fprintln(v, value)
	}

	v.Title = "Last Commits"

	return v
}

func MenuCursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if pos+2 > len(commitsIds) {
			// reatch the bottom of the list
			return nil
		}

		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
		pos++
	}
	return nil
}

func MenuCursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
		pos--
	}
	return nil
}
