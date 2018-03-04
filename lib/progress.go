// Display progress indicator
package lib

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
	todo := counter.total - counter.count
	perc := 100 - (todo / (counter.total / 100))
	if perc == counter.next {
		counter.next = perc + granularity
		fmt.Println(fmt.Sprintf("%d%% done", perc))
	}
}
