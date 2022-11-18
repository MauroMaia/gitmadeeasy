package main

import (
	"fmt"
	"os"

	"github.com/MauroMaia/gitmadeeasy/pkg/ui/branch"
	commit2 "github.com/MauroMaia/gitmadeeasy/pkg/ui/commit"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/help"
	menu "github.com/MauroMaia/gitmadeeasy/pkg/ui/menu"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

// These values may be set by the build script via the LDFLAGS argument
var (
	commit      string
	date        string
	version     string
	buildSource = "unknown"
)

func main() {
	utils.Logger.Infoln("##############\n")
	utils.Logger.Infof("# Version %s\n", version)
	utils.Logger.Infof("# Build Date %s\n", date)
	utils.Logger.Infof("# Commit Id %s\n", commit)
	utils.Logger.Infof("# Build Source %s\n", buildSource)
	utils.Logger.Infoln("##############\n")

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

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		utils.Logger.Panicln(err)
	}

	if err := menu.Keybindings(g); err != nil {
		utils.Logger.Panicln(err)
	}
	if err := branch.Keybindings(g); err != nil {
		utils.Logger.Panicln(err)
	}
	if err := commit2.Keybindings(g); err != nil {
		utils.Logger.Panicln(err)
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		utils.Logger.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	_, maxY := g.Size()

	menu.LayoutTopMenuOptions(g, -1, 0, maxY-3)

	_, _, xEnd, _, _ := g.ViewPosition(constants.MENU_VIEW)
	//log.Printf("xStart %d xEnd %d yStart %d yEnd %d", xStart,xEnd,yStart,yEnd)
	branch.LayoutListBranches(g, xEnd+1, 0)

	_, _, xEnd, _, _ = g.ViewPosition(constants.BRANCH_LIST_VIEW)
	//log.Printf("xStart %d xEnd %d yStart %d yEnd %d", xStart,xEnd,yStart,yEnd)
	commit2.LayoutListCommits(g, xEnd+1, 0)

	help.LayoutShowHelpView(g, -1, maxY-2)

	if _, err := utils.SetCurrentViewOnTop(g, constants.SELECTED_MENU); err != nil {
		utils.Logger.Fatalln(err)
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	_, _ = fmt.Fprintln(os.Stdout, "quiting")
	return gocui.ErrQuit
}
