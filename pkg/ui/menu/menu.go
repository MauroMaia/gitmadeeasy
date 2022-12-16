package ui

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/branch"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/push"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
	"strings"
)

const B_NEW_BRANCH = "New Branch"
const B_SHOW_COMMITS = "Commits List"
const B_SHOW_BRANCHS = "Branch's List"
const B_COMMIT_CHANGES = "Commit changes"
const B_PULL = "Pull"
const B_PUSH = "Push"

var B_PULL_IN_DYSPLAY = B_PULL
var B_PUSH_IN_DYSPLAY = B_PUSH

const B_SETTINGS = "Settings"

var buttons [10]string

func init() {
	createLabels()
}

func createLabels() {

	ahead, err := gitcmd.GetNrOfCommitsAhead()
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	behind, err := gitcmd.GetNrOfCommitsBehind()
	if err != nil {
		utils.Logger.Fatalln(err)
	}

	B_PULL_IN_DYSPLAY = B_PULL
	if behind != "0" {
		B_PULL_IN_DYSPLAY = B_PULL + " " +
			utils.TextToRed(behind+" ↓")
	}

	B_PUSH_IN_DYSPLAY = B_PUSH
	BPushToUi := B_PUSH
	if ahead != "0" {
		B_PUSH_IN_DYSPLAY = B_PUSH + " " + ahead + " ↑"
		BPushToUi = B_PUSH + " " +
			utils.TextToGreen(ahead+" ↑")
	}

	buttons = [10]string{
		B_NEW_BRANCH,
		B_SHOW_BRANCHS,
		"---",
		B_SHOW_COMMITS,
		B_COMMIT_CHANGES,
		"---",
		B_PULL_IN_DYSPLAY,
		BPushToUi,
		"---",
		B_SETTINGS,
	}
}

func LayoutTopMenuOptions(g *gocui.Gui, xBegins int, yBegins int, yEnd int) *gocui.View {

	// FIXME - move this to background task / timer
	createLabels()

	var stringLen = 0
	for _, str := range buttons {
		if len(str) > stringLen {
			stringLen = len(str)
		}
	}

	v, err := g.SetView(constants.MENU_VIEW, xBegins, yBegins, stringLen+1, yEnd)
	if err != nil && err != gocui.ErrUnknownView {
		utils.Logger.Fatalln(err)
	}

	v.Clear()

	for _, value := range buttons {
		_, _ = fmt.Fprintln(v, value)
	}

	v.Title = "MENU"
	v.Highlight = true

	return v
}

func MenuCursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy+2 > len(buttons) {
			// reatch the bottom of the list
			return nil
		}

		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
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

	switch strings.Trim(l, " ") {
	case B_NEW_BRANCH:
		branch.DisplayPopUp(g)
		constants.HIGHLIGHT_VIEW = constants.NEW_BRANCH_POPUP
		break
	case B_SHOW_COMMITS:
		g.DeleteView(constants.LEFT_VIEW)
		g.DeleteView(constants.RIGTH_VIEW)
		constants.LEFT_VIEW = constants.COMMIT_LIST_VIEW
		constants.RIGTH_VIEW = ""
		//TODO - constants.RIGTH_VIEW = constants.SHOW_CHANGES_IN_COMMIT_VIEW
		constants.HIGHLIGHT_VIEW = constants.COMMIT_LIST_VIEW
		break
	case B_SHOW_BRANCHS:
		g.DeleteView(constants.LEFT_VIEW)
		g.DeleteView(constants.RIGTH_VIEW)
		constants.LEFT_VIEW = constants.BRANCH_LIST_VIEW
		constants.HIGHLIGHT_VIEW = constants.BRANCH_LIST_VIEW
		break
	case B_COMMIT_CHANGES:
		g.DeleteView(constants.LEFT_VIEW)
		g.DeleteView(constants.RIGTH_VIEW)
		constants.LEFT_VIEW = constants.FILE_CHANGED_VIEW
		constants.RIGTH_VIEW = constants.DIFF_VIEW
		constants.HIGHLIGHT_VIEW = constants.FILE_CHANGED_VIEW
		break
	case B_PULL_IN_DYSPLAY:
		break
	case B_PUSH_IN_DYSPLAY:
		push.DisplayPopUp(g)
		constants.HIGHLIGHT_VIEW = constants.PUSH_POPUP
		break
	default:
		// not expected to do anything
		return nil
	}

	return nil
}

func Keybindings(g *gocui.Gui) error {

	if err := g.SetKeybinding(constants.MENU_VIEW, gocui.KeyArrowDown, gocui.ModNone, MenuCursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.MENU_VIEW, gocui.KeyArrowUp, gocui.ModNone, MenuCursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding(constants.MENU_VIEW, gocui.KeyEnter, gocui.ModNone, onEnterPress); err != nil {
		return err
	}
	return nil
}
