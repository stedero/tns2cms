package stats

import (
	"fmt"
)

// ProgressIndicator returns a function that displays progress in a percentage
// of the total given. It only displays the progress at every 10% step.
func ProgressIndicator(fileCount int) func() {
	granularity := 10
	total := fileCount
	count := 0
	next := granularity
	return func() {
		count += 1
		todo := float64(total - count)
		perc := int(100.0 - (todo / (float64(total) / 100.0)))
		if perc >= next {
			next += granularity
			fmt.Println(fmt.Sprintf("%d%% done", perc))
		}
	}
}
