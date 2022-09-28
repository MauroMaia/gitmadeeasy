package ui

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/jroimartin/gocui"
	"log"
)

func LayoutListBranches(g *gocui.Gui, xBegins int, yBegins int) {
	var branches = gitcmd.ListBranches()
	var stringLen = 0
	for _, branch := range branches {
		if len(branch.GetName()) > stringLen {
			stringLen = len(branch.GetName())
		}
	}

	_, maxY := g.Size()
	if len(branches) < maxY {
		maxY = len(branches)
	}
	if v, err := g.SetView("list_branches", xBegins, yBegins, stringLen, maxY+1); err != nil {
		if err != gocui.ErrUnknownView {
			log.Fatalln(err)
		}

		for _, value := range branches {
			var name = value.GetName()
			_, _ = fmt.Fprintln(v, name)
		}
	}
}
