package utils

import (
	"time"
)

type Timer struct {
	ch chan int64
}

func (t *Timer) Start() {
	if t.ch == nil {
		t.ch = make(chan int64, 1)
	}
	t.ch <- time.Now().UnixMilli()
}

func (t *Timer) End() int64 {
	return time.Now().UnixMilli() - <-t.ch
}

func (t *Timer) Close() {
	close(t.ch)
}
