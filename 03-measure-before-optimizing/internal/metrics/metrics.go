package metrics

import (
	"fmt"
	"sync/atomic"
)

// Collector is intentionally tiny.
//
// TODO: Replace or extend this with real counters and latency distributions.
// A useful next step is a request counter plus a histogram or summary-like
// approach for request duration.
type Collector struct {
	requests  atomic.Int64
	cacheHits atomic.Int64
}

func NewCollector() *Collector {
	return &Collector{}
}

func (c *Collector) ObserveRequest() {
	c.requests.Add(1)
}

func (c *Collector) ObserveCacheHit() {
	c.cacheHits.Add(1)
}

func (c *Collector) RenderPrometheus() string {
	return fmt.Sprintf(
		"# TODO: replace placeholder metrics\nrequests_total %d\ncache_hits_total %d\n",
		c.requests.Load(),
		c.cacheHits.Load(),
	)
}
