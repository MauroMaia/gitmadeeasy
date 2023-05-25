package ui

import (
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/branch"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/commit"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/log"
	status "github.com/MauroMaia/gitmadeeasy/pkg/ui/status"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

func DrawLeftView(g *gocui.Gui, xBegins int, painelXsize int, viewName string) {
	xEnd := xBegins + painelXsize - 1
	factory(g, xBegins+1, xEnd, 0, viewName)
}

func DrawRightView(g *gocui.Gui, xBegins int, painelXsize int, viewName string) {
	xEnd := xBegins + painelXsize
	factory(g, xBegins+1, xEnd, 0, viewName)
}

func factory(g *gocui.Gui, xBegin int, xEnd int, yBegin int, viewName string) {
	switch viewName {
	case constants.FILE_CHANGED_VIEW:
		status.LayoutShowStatus(g, xBegin, yBegin, xEnd)
		break
	case constants.BRANCH_LIST_VIEW:
		branch.LayoutListBranches(g, xBegin, yBegin, xEnd)
		break
	case constants.COMMIT_LIST_VIEW:
		commit.LayoutListCommits(g, xBegin, yBegin, xEnd)
		break
	case constants.DIFF_VIEW:
		commit.LayoutDiff(g, xBegin, yBegin, xEnd)
		break
	case constants.LOG_VIEW:
		log.Layout(g, xBegin, yBegin, xEnd)
		break
	default:
		utils.Logger.Warnf("This view has no process to be created on the left panel %s", viewName)
	}
}
