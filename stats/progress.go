package stats

import (
	"log"
	"time"

	"ibfd.org/tns2cms/paths"
)

// Reporter definition
type Reporter struct {
	start  time.Time
	counts []int
}

// NewReporter creates a reporter.
func NewReporter() *Reporter {
	return &Reporter{time.Now(), make([]int, paths.NumActions)}
}

// End reports the termination of the process with counts.
func (reporter *Reporter) End() {
	elapsed := time.Since(reporter.start)
	log.Printf("skipped %d directories", reporter.counts[paths.RejectDir])
	log.Printf("skipped %d files", reporter.counts[paths.RejectFile])
	log.Printf("processing %d files in %d directories took %s", reporter.counts[paths.AcceptFile], reporter.counts[paths.AcceptDir], elapsed)
}

// Register an action.
func (reporter *Reporter) Register(action paths.Action, path string) {
	reporter.counts[action]++
	switch action {
	case paths.RejectDir:
		log.Printf("skipped directory: %s\n", path)
	case paths.RejectFile:
		log.Printf("skipped file: %s\n", path)
	default:
	}
}
