package count

import (
	"strings"
)

func Count(str string) int {
	items := strings.Split(str, " ")
	return len(items)
}
