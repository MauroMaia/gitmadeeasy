package ui

import (
	"fmt"
	"log"

	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/jroimartin/gocui"
)

const COMMIT_LIST = "COMMIT_LIST"

func LayoutListCommits(g *gocui.Gui, xBegins int, yBegins int) *gocui.View {

	var commitsIds = gitcmd.ListCommitIDs()
	var stringLen = len(commitsIds[0])

	_, maxY := g.Size()
	if len(commitsIds) < maxY {
		maxY = len(commitsIds)
	}

	v, err := g.SetView(COMMIT_LIST, xBegins, yBegins, stringLen+1, yBegins+maxY+1)
	if err != nil && err != gocui.ErrUnknownView {
		log.Fatalln(err)
	}

	v.Clear()

	for _, value := range commitsIds {
		_, _ = fmt.Fprintln(v, value)
	}

	v.Title = "Last Commits"

	return v
}
