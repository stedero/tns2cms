package stats

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/inflection"
	"ibfd.org/tns2cms/paths"
)

// Reporter definition
type Reporter struct {
	start   time.Time
	verbose bool
	counts  []int
}

// NewReporter creates a reporter.
func NewReporter(verbose bool) *Reporter {
	return &Reporter{time.Now(), verbose, make([]int, paths.NumActions)}
}

// Register an action.
func (reporter *Reporter) Register(action paths.Action, path string) {
	reporter.counts[action]++
	if reporter.verbose {
		switch action {
		case paths.RejectDir:
			log.Printf("skipped directory: %s\n", path)
		case paths.RejectFile:
			log.Printf("skipped file: %s\n", path)
		default:
		}
	}
}

// End logs the termination of the process with counts.
func (reporter *Reporter) End() {
	counts := reporter.counts
	elapsed := time.Since(reporter.start)
	dirFmt := countFmt("directory")
	fileFmt := countFmt("file")
	log.Printf("skipped %s", dirFmt(counts[paths.RejectDir]))
	log.Printf("skipped %s", fileFmt(counts[paths.RejectFile]))
	log.Printf("processing %s in %s took %s", fileFmt(counts[paths.AcceptFile]), dirFmt(counts[paths.AcceptDir]), elapsed)
}

func countFmt(singular string) func(int) string {
	return func(count int) string {
		return fmtPluralize(singular, count)
	}
}

func fmtPluralize(singular string, count int) string {
	return fmt.Sprintf("%d %s", count, pluralize(singular, count))
}

func pluralize(singular string, count int) string {
	if count == 1 {
		return singular
	}
	return inflection.Plural(singular)
}
