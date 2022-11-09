package ui

import (
	"fmt"
	"log"

	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
)

const TOP_MENU = "top_menu"
var buttons = [2]string{"New Branch","Commit"}

func LayoutTopMenuOptions(g *gocui.Gui, xBegins int, yBegins int,yEnd int) *gocui.View {
	var stringLen = 0
	for _, str := range buttons {
		if len(str) > stringLen {
			stringLen += len(str)
		}
	}

	v, err := g.SetView(TOP_MENU, xBegins, yBegins, stringLen+1, yEnd)

	if err != nil && err != gocui.ErrUnknownView {
		log.Fatalln(err)
	}

	for _, value := range buttons {
		var name = value
		_, _ = fmt.Fprintln(v, name)
	}

	v.Title="Cmds"

	if _, err = utils.SetCurrentViewOnTop(g, TOP_MENU); err != nil {
		log.Fatalln(err)
	}
	
	return v
}


func buttonView(name string,v *gocui.View){
	//stringLen := len(name)
	//_, _ = fmt.Fprintln(v, name)
}