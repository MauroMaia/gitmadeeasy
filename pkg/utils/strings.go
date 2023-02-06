package utils

import "fmt"

func DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func TextToRed(in string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m ", 1, in)
}

func TextToGreen(in string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m ", 40, in)
}

func TextToYellow(in string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m ", 11, in)
}

func TextToOrange(in string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m ", 172, in)
}

func TextToBlueWater(in string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m ", 45, in)
}

func TextToBlue(in string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m ", 21, in)
}
