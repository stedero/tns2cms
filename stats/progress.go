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
	done, next := 0, 10
	return func() {
		done++
		perc := int(float64(done) / onePercent)
		if perc >= next {
			next += 10
			fmt.Fprintln(os.Stderr, fmt.Sprintf("%d%% done", perc))
		}
	}
}

// Reporter returns a function that displays the time it took to
// process a number of files.
func Reporter(fileCount int) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		fmt.Fprintf(os.Stderr, "processing %d files took %s\n", fileCount, elapsed)
	}
}
