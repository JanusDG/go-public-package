package lib

import (
	"fmt"
	"time"
)

type Timer struct {
	startPoint time.Time
	endPoint   time.Time
	flags      map[int]time.Time
	flagCount  int
}

func NewTimer() *Timer {
	return &Timer{flags: make(map[int]time.Time)}
}

func (t *Timer) StartTimer() {
	t.startPoint = time.Now()
}

func (t *Timer) EndTimer() {
	t.endPoint = time.Now()
}

func (t *Timer) FlagNow() {
	t.flags[t.flagCount] = time.Now()
	t.flagCount++
}

func (t Timer) Show() {
	for k, v := range t.flags {
		fmt.Println(k, v)
	}
	fmt.Println("\nTotal time: ", t.endPoint.Sub(t.startPoint))
}
