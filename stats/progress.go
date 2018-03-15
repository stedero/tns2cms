package stats

import (
	"fmt"
	"os"
	"time"
	"tns2cms/naming"
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
func (reporter *Reporter) Register(validation naming.Validation, path string) {
	switch validation {
	case naming.AcceptDir:
		reporter.directoriesAccepted++
	case naming.AcceptFile:
		reporter.filesAccepted++
	case naming.RejectDir:
		fmt.Fprintf(os.Stderr, "skipped directory: %s\n", path)
	case naming.RejectFile:
		fmt.Fprintf(os.Stderr, "skipped file: %s\n", path)
	}
}
