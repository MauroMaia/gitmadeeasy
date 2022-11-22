package main

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/help"
	"os"

	"github.com/MauroMaia/gitmadeeasy/pkg/ui/branch"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/commit"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	menu "github.com/MauroMaia/gitmadeeasy/pkg/ui/menu"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

// These values may be set by the build script via the LDFLAGS argument
var (
	commitHash  string
	date        string
	version     string
	buildSource = "unknown"
)

func main() {
	utils.Logger.Infoln("##############")
	utils.Logger.Infof("# Version %s\n", version)
	utils.Logger.Infof("# Build Date %s\n", date)
	utils.Logger.Infof("# Commit Id %s\n", commitHash)
	utils.Logger.Infof("# Build Source %s\n", buildSource)
	utils.Logger.Infoln("##############")

	if !utils.IsGitRepoDirectory() {
		utils.Logger.Fatalln("Directory .git not found")
	}

	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		utils.Logger.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	setKeybindings(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		utils.Logger.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	// get windows size
	maxX, maxY := g.Size()

	//
	//	DEFAULT UI
	//
	help.LayoutShowHelpView(g, -1, maxY-2)
	menu.LayoutTopMenuOptions(g, -1, 0, maxY-3)

	_, _, xEnd, _, _ := g.ViewPosition(constants.MENU_VIEW)
	painelXsize := (maxX - xEnd) / 2

	//
	// PANELS
	//

	ui.DrawLeftView(g, xEnd, painelXsize, constants.LEFT_VIEW)

	if constants.RIGTH_VIEW != "" {
		_, _, xEnd, _, _ = g.ViewPosition(constants.LEFT_VIEW)
		ui.DrawRightView(g, xEnd, painelXsize, constants.RIGTH_VIEW)
	}

	if _, err := utils.SetCurrentViewOnTop(g, constants.HIGHLIGTH_VIEW); err != nil {
		utils.Logger.Fatalln(err)
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	_, _ = fmt.Fprintln(os.Stdout, "quiting")
	return gocui.ErrQuit
}

func setKeybindings(g *gocui.Gui) {

	//
	//	Global Keybindings
	//
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		utils.Logger.Panicln(err)
	}

	//
	// Panel specific Keybindings
	//
	if err := menu.Keybindings(g); err != nil {
		utils.Logger.Panicln(err)
	}
	if err := branch.Keybindings(g); err != nil {
		utils.Logger.Panicln(err)
	}
	if err := commit.Keybindings(g); err != nil {
		utils.Logger.Panicln(err)
	}
}
