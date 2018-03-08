package stats

import (
	"fmt"
	"os"
)

// ProgressIndicator returns a function that displays progress
// on stderr as a percentage of the total given.
func ProgressIndicator(total int) func() {
	onePercent := float64(total) / 100.0
	done, next := 0, 1
	return func() {
		done += 1
		perc := int(float64(done) / onePercent)
		if perc >= next {
			next += 1
			fmt.Fprintln(os.Stderr, fmt.Sprintf("%d%% done", perc))
		}
	}
}
