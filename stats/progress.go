package stats

import (
	"fmt"
	"os"
)

// ProgressIndicator returns a function that displays progress
// as a percentage of the total given on stderr.
func ProgressIndicator(total int) func() {
	onePercent := float64(total) / 100.0
	done, next := 0, 1
	return func() {
		done++
		perc := int(float64(done) / onePercent)
		if perc >= next {
			next++
			fmt.Fprintln(os.Stderr, fmt.Sprintf("%d%% done", perc))
		}
	}
}
