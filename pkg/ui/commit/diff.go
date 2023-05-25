package commit

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

var diffLines []string
var pos = 0
var maxPos = 0

func init() {
	diffLines, _ = gitcmd.GetDiffPatch()
	maxPos = len(diffLines)
}

// TODO - fill the docs
func SetDiffForFile(filename string) {
	diffLines, _ = gitcmd.GetDiffPatchForFile(filename)
	maxPos = len(diffLines)
}

func LayoutDiff(g *gocui.Gui, xBegins int, yBegins int, xEnd int) *gocui.View {

	_, maxY := g.Size()

	v, err := g.SetView(constants.DIFF_VIEW, xBegins, yBegins, xEnd, maxY-2)
	if err != nil && err != gocui.ErrUnknownView {
		utils.Logger.Fatalln(err)
	}

	v.Clear()

	for _, value := range diffLines {
		if len(value) > 0 && value[0:1] != " " {
			switch value[0:1] {
			case "+":
				_, _ = fmt.Fprintln(v, utils.TextToGreen(value))
			case "-":
				_, _ = fmt.Fprintln(v, utils.TextToRed(value))
			default:
				_, _ = fmt.Fprintln(v, value)
			}
		} else {
			_, _ = fmt.Fprintln(v, value)
		}
	}

	v.Title = "Diff"
	// TODO - create an option in Settings
	v.Wrap = true

	return v
}

// TODO - fill the docs
func diffCursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if pos+2 > maxPos {
			// reach the bottom of the list
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

// TODO - fill the docs
func diffCursorUp(g *gocui.Gui, v *gocui.View) error {
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
