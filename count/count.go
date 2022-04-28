package count

import (
	"strings"
)

func Count(str string) int {
	if str == "" {
		return 0
	}
	items := strings.Split(str, " ")
	return len(items)
}
