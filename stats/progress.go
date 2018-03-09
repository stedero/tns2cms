package stats

import (
	"fmt"
	"os"
	"time"
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

// Reporter returns a function that displays the elapsed time for an action.
// The text is formatted using the template and the count parameters as
// input to fmt.Sprintf(template, count). So the template must contain a
// %d placeholder for the count.
func Reporter(template string, count int) func() {
	start := time.Now()
	action := fmt.Sprintf(template, count)
	return func() {
		elapsed := time.Since(start)
		fmt.Fprintf(os.Stderr, "%s took %s\n", action, elapsed)
	}
}
