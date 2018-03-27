package stats

import (
	"fmt"
	"github.com/jinzhu/inflection"
	"ibfd.org/tns2cms/paths"
	"log"
	"time"
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

// End reports the termination of the process with counts.
func (reporter *Reporter) End() {
	elapsed := time.Since(reporter.start)
	dirFmt := countFmt("directory")
	fileFmt := countFmt("file")
	log.Printf("skipped %s", dirFmt(reporter.counts[paths.RejectDir]))
	log.Printf("skipped %s", fileFmt(reporter.counts[paths.RejectFile]))
	log.Printf("processing %s in %s took %s", fileFmt(reporter.counts[paths.AcceptFile]), dirFmt(reporter.counts[paths.AcceptDir]), elapsed)
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
	} else {
		return inflection.Plural(singular)
	}
}
