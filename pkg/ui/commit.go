package ui

import (
	"fmt"
	"log"

	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/jroimartin/gocui"
)

const COMMIT_LIST = "COMMIT_LIST"

func LayoutListCommits(g *gocui.Gui, xBegins int, yBegins int) *gocui.View {

	//log.Fatalf("xBegins %d yBegins%d \n",xBegins,yBegins)

	var commitsIds = gitcmd.ListCommitIDs()
	var stringLen = len(commitsIds[0])

	_, maxY := g.Size()
	//log.Fatalf("maxX %d stringLen %d xBegins %d\n",maxX,stringLen,xBegins)

	v, err := g.SetView(COMMIT_LIST, xBegins, yBegins, xBegins+stringLen+2, maxY-4)
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
