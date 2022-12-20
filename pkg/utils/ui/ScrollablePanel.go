package ui

import (
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

type ScrollablePanel struct {
	name           string
	title          string
	content        string
	x, y           int
	w, h           int
	cursorLine     int
	wrap, editable bool
}

func NewScrollablePanel(name string, x int, y int, w int, h int) *ScrollablePanel {
	return &ScrollablePanel{
		name:       name,
		title:      "",
		x:          x,
		y:          y,
		w:          w,
		h:          h,
		cursorLine: 0,
		wrap:       true,
		editable:   false,
	}
}

func (s *ScrollablePanel) SetTitle(title string) *ScrollablePanel {
	s.title = title
	return s
}
func (s *ScrollablePanel) SetContent(content string) *ScrollablePanel {
	s.content = content
	return s
}

func (s *ScrollablePanel) Layout(g *gocui.Gui) error {
	v, err := g.SetView(s.name, s.x, s.y, s.x+s.w, s.y+s.h)
	if err != nil && err != gocui.ErrUnknownView {
		utils.Logger.Fatalln(err)
	}

	v.Title = s.title
	v.Wrap = s.wrap
	v.Editable = s.editable

	v.Clear()
	_, err = fmt.Fprint(v, "PANEL IN TEST")
	_, err = fmt.Fprint(v, s.content)

	if err != nil {
		utils.Logger.Fatalln(err)
	}

	return nil
}
