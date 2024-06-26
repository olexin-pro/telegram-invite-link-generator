package limiter

import (
	"sync"
	"time"
)

type Limiter interface {
	Wait()
}

type limiter struct {
	tick    time.Duration
	count   uint
	entries []time.Time
	index   uint
	mutex   sync.Mutex
}

func NewLimiter(tick time.Duration, count uint) Limiter {
	l := limiter{
		tick:  tick,
		count: count,
		index: 0,
	}
	l.entries = make([]time.Time, count)
	before := time.Now().Add(-2 * tick)
	for i, _ := range l.entries {
		l.entries[i] = before
	}
	return &l
}
func (l *limiter) Wait() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	last := &l.entries[l.index]
	next := last.Add(l.tick)
	now := time.Now()
	if now.Before(next) {
		time.Sleep(next.Sub(now))
	}
	*last = time.Now()
	l.index = l.index + 1
	if l.index == l.count {
		l.index = 0
	}
}
