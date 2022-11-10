package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MauroMaia/gitmadeeasy/pkg/ui"
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
	log.Printf("##############\n")
	log.Printf("# Version %s\n", version)
	log.Printf("# Build Date %s\n", date)
	log.Printf("# Commit Id %s\n", commit)
	log.Printf("# Build Source %s\n", buildSource)
	log.Printf("##############\n")

	if !utils.IsGitRepoDirectory() {
		log.Fatalln("Directory .git not found")
	}

	g,err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := menu.Keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	_,maxY := g.Size()

	menu.LayoutTopMenuOptions(g, -1, 0, maxY-1)

	_, _, xEnd, _, _ := g.ViewPosition(menu.TOP_MENU)
	//xStart, yStart, xEnd, yEnd, _ := g.ViewPosition(menu.TOP_MENU)
	//log.Printf("xStart %d xEnd %d yStart %d yEnd %d", xStart,xEnd,yStart,yEnd)
	ui.LayoutListBranches(g, xEnd+1, 0)

	_, _, xEnd, _, _  = g.ViewPosition(ui.BRANCH_LIST)
	//xStart, yStart, xEnd, yEnd, _ := g.ViewPosition(ui.BRANCH_LIST)
	//log.Printf("xStart %d xEnd %d yStart %d yEnd %d", xStart,xEnd,yStart,yEnd)
	ui.LayoutListCommits(g, xEnd+1, 0)

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	fmt.Fprintln(os.Stdout, "quiting")
	return gocui.ErrQuit
}
