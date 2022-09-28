package main

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui"
	"github.com/jroimartin/gocui"
	"log"
	"os"
)

func main() {
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

	ui.LayoutListBranches(g, 0, 0)

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	fmt.Fprintln(os.Stdout, "quiting")
	return gocui.ErrQuit
}
