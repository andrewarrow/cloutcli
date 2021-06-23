package display

import (
	"fmt"
	"strings"
)

func Float(val float64) string {
	return fmt.Sprintf("%.02f", val)
}
func OneE9(val int64) string {
	return fmt.Sprintf("%.02f", OneE9Float(val))
}
func OneE9extra(val int64) string {
	return fmt.Sprintf("%.06f", OneE9Float(val))
}
func OneE9Float(val int64) float64 {
	return float64(val) / 1000000000.0
}

func Header(sizes []int, fields ...string) {
	for i, field := range fields {
		fmt.Printf("%s ", LeftAligned(field, sizes[i]))
	}
	fmt.Printf("\n")
	for i, field := range fields {
		dashes := []string{}
		for i := 0; i < len(field); i++ {
			dashes = append(dashes, "-")
		}
		fmt.Printf("%s ", LeftAligned(strings.Join(dashes, ""), sizes[i]))
	}
	fmt.Printf("\n")
}
func Row(sizes []int, items ...interface{}) {
	for i, item := range items {
		fmt.Printf("%s ", LeftAligned(item, sizes[i]))
	}
	fmt.Printf("\n")
}

func LeftAligned(thing interface{}, size int) string {
	s := fmt.Sprintf("%v", thing)

	if len(s) >= size {
		return s[0:size-1] + " "
	}
	fill := size - len(s)
	spaces := []string{}
	for {
		spaces = append(spaces, " ")
		if len(spaces) >= fill {
			break
		}
	}
	return s + strings.Join(spaces, "")
}
