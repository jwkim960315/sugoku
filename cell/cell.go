package cell

import (
	"strconv"
)

func GenerateCellContent(number int) string {
	return " " + strconv.Itoa(number) + " "
}