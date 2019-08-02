package subscriber

import (
	"fmt"
	"time"
)

type color string

var (
	red    color = "#e23737"
	green  color = "#14a127"
	yellow color = "#faef56"
)

func withColor(str string, c color) string {
	return fmt.Sprintf("<span style='color:%s'>%s</span>", c, str)
}

func withTime(str string) string {
	return fmt.Sprintf("%s> %s", withColor(time.Now().Format("00:00:00.0000"), green), str)
}
