package q2

import (
	"fmt"
	"sort"
	"strings"
)

func Run(input []int) string {
	sort.Ints(input)
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(input)), " "), "[]")
}
