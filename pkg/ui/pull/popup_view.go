package pull

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils/ui"
	"github.com/jroimartin/gocui"
	"regexp"
	"strings"
)

var p ui.OptionsPanel

var options []interface{}

func init() {
	prune := ui.NewOptionsElement[bool](
		"Prune",
		"delete old branches",
		false,
	)
	stash := ui.NewOptionsElement[bool](
		"Stash",
		"save/apply local changes after pulling",
		false,
	)
	rebase := ui.NewOptionsElement[bool](
		"Rebase",
		"apply remote changes before local commits",
		false,
	)
	options = []interface{}{
		stash,
		prune,
		rebase,
	}
}

var optionLineRegex, _ = regexp.Compile("^\\[...(.*)")

func DisplayPopUp(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	p = ui.NewOptionsPanel(constants.PUll_POPUP,
		"PULL",
		maxX/2-25,
		maxY/2-len(options),
		50,
		&options,
	)
	p.Layout(g)

	if _, err := g.SetCurrentView(constants.PUll_POPUP); err != nil {
		return err
	}

	if _, err := utils.SetCurrentViewOnTop(g, constants.PUll_POPUP); err != nil {
		utils.Logger.Fatalln(err)
	}

	constants.HIGHLIGHT_VIEW = constants.PUll_POPUP

	return nil
}

func onEnterPress(g *gocui.Gui, v *gocui.View) error {

	v.Clear()
	X, _ := v.Size()

	rep := int(0.5 * float64(X-1))
	fmt.Fprint(v, strings.Repeat("▒", rep))
	fmt.Fprint(v, " 50%")

	go func() {
		_, err := gitcmd.Pull(options[1].(ui.OptionsElement[bool]).Val, options[2].(ui.OptionsElement[bool]).Val, false)

		g.Update(func(g *gocui.Gui) error {
			if err != nil {
				v.Clear()
				v.Editable = false
				v.BgColor = gocui.ColorRed
				v.Rewind()
				_, _ = fmt.Fprintln(v, err.Error())
				v.SetCursor(0, 0)
			} else {
				g.DeleteView(constants.PUll_POPUP)
				constants.LEFT_VIEW = constants.COMMIT_LIST_VIEW
				constants.RIGTH_VIEW = constants.BRANCH_LIST_VIEW
				constants.HIGHLIGHT_VIEW = constants.MENU_VIEW
			}
			return nil
		})
	}()

	return nil
}

func onEnterSpace(g *gocui.Gui, v *gocui.View) error {

	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}
	s := optionLineRegex.FindStringSubmatch(l)
	s1 := s[1]

	if l != "" && len(s) > 0 {
		switch s1 {
		case "Stash":
			optionConv := options[0].(ui.OptionsElement[bool])
			optionConv.Val = !optionConv.Val
			options[0] = optionConv
			break
		case "Prune":
			optionConv := options[1].(ui.OptionsElement[bool])
			optionConv.Val = !optionConv.Val
			options[1] = optionConv
			break
		case "Rebase":
			optionConv := options[2].(ui.OptionsElement[bool])
			optionConv.Val = !optionConv.Val
			options[2] = optionConv
			break
		}
	}

	g.Update(func(gui *gocui.Gui) error {
		return p.Update(gui, &options)
	})

	return nil
}

// TODO - fill the docs
func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy+2 >= len(options)*2 { // FIXME - remove this magical number
			// reatch the bottom of the list
			return nil
		}

		if err := v.SetCursor(cx, cy+2); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+2); err != nil {
				return err
			}
		}
	}
	return nil
}

// TODO - fill the docs
func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy-2 < 0 { // FIXME - remove this magical number
			// reatch the bottom of the list
			return nil
		}

		if err := v.SetCursor(cx, cy-2); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy-2); err != nil {
				return err
			}
		}
	}
	return nil
}

func quitPopup(g *gocui.Gui, v *gocui.View) error {

	g.DeleteView(constants.PUll_POPUP)
	constants.HIGHLIGHT_VIEW = constants.MENU_VIEW

	return nil
}
