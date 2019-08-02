package subscriber

import (
	"fmt"
	"time"
)

type color string

var (
	red    color = "red"
	green  color = "green"
	yellow color = "yellow"
)

func withColor(str string, c color) string {
	return fmt.Sprintf("<span class='%s'>%s</span>", c, str)
}

func withTime(str string) string {
	return fmt.Sprintf("%s> %s", withColor(time.Now().Format("15:04:05.000"), green), str)
}
