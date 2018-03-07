// Display progress indicator
package stats

import (
	"fmt"
)

const granularity = 10

type Counter struct {
	total int
	count int
	next  int
}

func NewCounter(total int) *Counter {
	return &Counter{total, 0, granularity}
}

func (counter *Counter) Next() {
	counter.count += 1
	todo := float64(counter.total - counter.count)
	perc := int(100.0 - (todo / (float64(counter.total) / 100.0)))
	if perc >= counter.next {
		counter.next += granularity
		fmt.Println(fmt.Sprintf("%d%% done", perc))
	}
}
