package branch

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/jroimartin/gocui"
	"log"
)

func LayoutListBranches(g *gocui.Gui, xBegins int, yBegins int) *gocui.View {
	var branches = gitcmd.ListBranches()
	var stringLen = 0
	for _, branch := range branches {
		if len(branch.GetName()) > stringLen {
			stringLen = len(branch.GetName())
		}
	}

	_, maxY := g.Size()

	v, err := g.SetView(constants.BRANCH_LIST_VIEW, xBegins, yBegins, xBegins+stringLen+2, maxY-4)
	if err != nil && err != gocui.ErrUnknownView {
		log.Fatalln(err)
	}

	v.Clear()

	for _, value := range branches {
		var name = value.GetName()
		_, _ = fmt.Fprintln(v, name)
	}

	v.Title = "Branch list"

	return v
}
