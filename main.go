package main

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
	"log"
	"os"
)

// These values may be set by the build script via the LDFLAGS argument
var (
	commit      string
	date        string
	version     string
	buildSource = "unknown"
)

func main() {
	if !utils.IsGitRepoDirectory() {
		log.Fatalln(os.Stderr, "Directory .git not found")
	}

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	gitcmd.ListCommitIDs()
	ui.LayoutListBranches(g, 0, 0)

	xBegins, _, _, yBegins, _ := g.ViewPosition(ui.BRANCH_LIST)

	ui.LayoutListCommits(g, xBegins, yBegins)

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	fmt.Fprintln(os.Stdout, "quiting")
	return gocui.ErrQuit
}
