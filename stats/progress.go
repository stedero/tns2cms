package stats

import (
	"fmt"
	"os"
	"time"
	"tns2cms/paths"
)

// Reporter definition
type Reporter struct {
	start               time.Time
	directoriesAccepted int
	filesAccepted       int
}

// NewReporter creates a reporter.
func NewReporter() *Reporter {
	return &Reporter{time.Now(), 0, 0}
}

// End reports the termination of the process with counts.
func (reporter *Reporter) End() {
	elapsed := time.Since(reporter.start)
	fmt.Fprintf(os.Stderr, "processing %d files in %d directories took %s\n", reporter.filesAccepted, reporter.directoriesAccepted, elapsed)
}

// Register a validation entry.
func (reporter *Reporter) Register(validation paths.Validation, path string) {
	switch validation {
	case paths.AcceptDir:
		reporter.directoriesAccepted++
	case paths.AcceptFile:
		reporter.filesAccepted++
	case paths.RejectDir:
		fmt.Fprintf(os.Stderr, "skipped directory: %s\n", path)
	case paths.RejectFile:
		fmt.Fprintf(os.Stderr, "skipped file: %s\n", path)
	}
}
