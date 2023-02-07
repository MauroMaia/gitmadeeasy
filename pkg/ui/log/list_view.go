package log

import (
	"encoding/json"
	"fmt"
	"github.com/MauroMaia/gitmadeeasy/pkg/ui/constants"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"github.com/jroimartin/gocui"
	"os"
	"strings"
)

var logLines []string
var posListView = 0

func Layout(g *gocui.Gui, xBegins int, yBegins int, xEnd int) *gocui.View {

	_, maxY := g.Size()

	v, err := g.SetView(constants.LOG_VIEW, xBegins, yBegins, xEnd, maxY-3)
	if err != nil && err != gocui.ErrUnknownView {
		utils.Logger.Fatalln(err)
	}

	v.Clear()

	logLines = getFormattedLog()
	for _, val := range logLines {
		_, _ = fmt.Fprintln(v, val)
	}

	v.Title = "Log"

	// TODO - create an option in Settings
	v.Wrap = false
	v.Editable = false

	return v
}

func menuCursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if posListView+2 > len(logLines) {
			// reach the bottom of the list
			return nil
		}

		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err = v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
		posListView++
	}
	return nil
}

func menuCursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err = v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
		posListView--
	}
	return nil
}

func getFormattedLog() []string {
	b, err := os.ReadFile(utils.FILE_PATH)

	if err != nil {
		utils.Logger.Fatal(err)
	}

	rawLogLines := strings.Split(string(b), "\n")

	var result []string
	for _, value := range rawLogLines {

		if strings.Trim(value, " ") == "" {
			continue
		}

		var objmap map[string]interface{}
		if err = json.Unmarshal([]byte(value), &objmap); err != nil {
			panic(err)
		}

		switch objmap["level"] {
		case "trace":
			var data = objmap["data"].(map[string]interface{})
			if data != nil {
				if val, ok := data["cmd"]; ok {
					//do something here
					result = append(result, fmt.Sprintf("%s - %6s - %s", objmap["time"], objmap["level"], objmap["msg"]))

					if data["status_code"] != 0.0 {
						result = append(result, fmt.Sprintf("[%3.0f] %s", data["status_code"].(float64), utils.TextToRed(val.(string))))
					} else {
						result = append(result, fmt.Sprintf("> %s", utils.TextToBlueWater(val.(string))))
					}

					for _, content := range data["output"].([]interface{}) {
						result = append(result, fmt.Sprintf(" %s", utils.TextToOrange(content.(string))))
					}

					break
				}
			}
			result = append(result, fmt.Sprintf("%s - %6s - %s", objmap["time"], objmap["level"], objmap["msg"]))
		default:
			result = append(result, fmt.Sprintf("%s - %6s - %s", objmap["time"], objmap["level"], objmap["msg"]))
		}
	}
	return result
}
