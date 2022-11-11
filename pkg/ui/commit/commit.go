package commit

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"log"

	"github.com/MauroMaia/gitmadeeasy/pkg/gitcmd"
	"github.com/jroimartin/gocui"
)

func LayoutListCommits(g *gocui.Gui, xBegins int, yBegins int) *gocui.View {

	//log.Fatalf("xBegins %d yBegins%d \n",xBegins,yBegins)

	var commitsIds = gitcmd.ListCommitIDs()
	var stringLen = len(commitsIds[0])

	_, maxY := g.Size()
	//log.Fatalf("maxX %d stringLen %d xBegins %d\n",maxX,stringLen,xBegins)

	v, err := g.SetView(constants.COMMIT_LIST_VIEW, xBegins, yBegins, xBegins+stringLen+2, maxY-3)
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
