package ui

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/jroimartin/gocui"
	"log"
)

const BRANCH_LIST = "list_branches"

func LayoutListBranches(g *gocui.Gui, xBegins int, yBegins int) *gocui.View {
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

	v, err := g.SetView(BRANCH_LIST, xBegins, yBegins, stringLen+1, maxY+1)

	if err != nil && err != gocui.ErrUnknownView {
		log.Fatalln(err)
	}

	for _, value := range branches {
		var name = value.GetName()
		_, _ = fmt.Fprintln(v, name)
	}
	return v
}
