package stats

import (
	"fmt"
)

// ProgressIndicator returns a function that displays progress as a percentage
// of the total given. It only displays the progress at every 10% step.
func ProgressIndicator(total int) func() {
	granularity := 10
	onePercent := float64(total) / 100.0
	done := 0
	next := granularity
	return func() {
		done += 1
		perc := int(float64(done) / onePercent)
		if perc >= next {
			next += granularity
			fmt.Println(fmt.Sprintf("%d%% done", perc))
		}
	}
}
