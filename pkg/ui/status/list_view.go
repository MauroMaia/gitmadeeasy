package commit

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/commit"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
	"strings"
	"time"
)

var content = ""
var pos = 0
var maxPos = 0

var ticker *time.Ticker
var done = make(chan bool)

func init() {
	utils.Logger.Traceln("init - status.list_view at")
	loadFilesStatusStatus()
	rescheduleTimer()
}

func loadFilesStatusStatus() {
	result := gitcmd.ListFilesChanged()
	result = utils.DeleteEmpty(result)
	maxPos = len(result)
	content = ""

	for _, value := range result {
		if len(value) > 0 && value[0:1] != " " {
			switch value[0:1] {
			case "M":
				content += utils.TextToYellow(value) + "\n"
			case "R":
				content += utils.TextToBlue(value) + "\n"
			case "A":
				content += utils.TextToGreen(value) + "\n"
			case "D":
				content += utils.TextToRed(value) + "\n"
			default:
				content += value + "\n"
			}
		} else {
			content += value + "\n"
		}
	}
}

func rescheduleTimer() {
	utils.Logger.Infoln("Ticker reschedule")

	ticker = time.NewTicker(15 * time.Second)

	go func() {
		for {
			select {
			case <-done:
				// Stops this ticker it's possible but not expected
				return
			case _ = <-ticker.C:
				utils.Logger.Debugln("Tick at")
				loadFilesStatusStatus()
			}
		}
	}()
}

func LayoutShowStatus(g *gocui.Gui, xBegins int, yBegins int, xEnd int) *gocui.View {

	_, maxY := g.Size()

	v, err := g.SetView(constants.FILE_CHANGED_VIEW, xBegins, yBegins, xEnd, maxY-3)
	if err != nil && err != gocui.ErrUnknownView {
		utils.Logger.Fatalln(err)
	}

	v.Clear()

	_, _ = fmt.Fprintln(v, content)

	v.Title = "Files Changed"

	return v
}

func MenuCursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if pos+2 > maxPos {
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

		var line string
		var err error

		if line, err = v.Line(cy + 1); err != nil || line == "" {
			return nil
		}
		commit.SetDiffForFile(strings.Trim(line[3:], " "))
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

		var line string
		var err error

		if line, err = v.Line(cy - 1); err != nil {
			return nil
		}
		commit.SetDiffForFile(strings.Trim(line[3:], " "))
	}
	return nil
}

func StageFile(g *gocui.Gui, v *gocui.View) error {
	var line string
	var err error

	_, cy := v.Cursor()
	if line, err = v.Line(cy); err != nil {
		return nil
	}

	gitcmd.StageFile(strings.Trim(line[3:], " "))
	loadFilesStatusStatus()

	return nil
}
