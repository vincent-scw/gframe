package subscriber

import "fmt"

type color string

var (
	red color    = "#e23737"
	green color  = "#14a127"
	yellow color = "#faef56"
)

func withColor(str string, c color) string {
	return fmt.Sprintf("<span style='color:%s'>%s</span>", c, str)
}