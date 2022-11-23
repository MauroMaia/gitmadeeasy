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
	commitsIds = gitcmd.ListFilesChanged()
}

func LayoutShowStatus(g *gocui.Gui, xBegins int, yBegins int, xEnd int) *gocui.View {

	_, maxY := g.Size()

	v, err := g.SetView(constants.FILE_CHANGED_VIEW, xBegins, yBegins, xEnd, maxY-3)
	if err != nil && err != gocui.ErrUnknownView {
		utils.Logger.Fatalln(err)
	}

	v.Clear()

	// FIXME - do this after data received to run only once
	for _, value := range commitsIds {
		if len(value) > 0 && value[0:1] != " " {
			switch value[0:1] {
			case "M":
				_, _ = fmt.Fprintln(v, utils.TextToYellow(value))
			case "R":
				_, _ = fmt.Fprintln(v, utils.TextToBlue(value))
			case "A":
				_, _ = fmt.Fprintln(v, utils.TextToGreen(value))
			case "D":
				_, _ = fmt.Fprintln(v, utils.TextToRed(value))
			default:
				_, _ = fmt.Fprintln(v, value)
			}
		} else {
			_, _ = fmt.Fprintln(v, value)
		}
	}

	v.Title = "Files Changed"

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
