// Package libTimer provides simple timer with flags
package libTimer

import (
	"fmt"
	"time"
)

// Timer implemets timing and flagging
type Timer struct {
	startPoint time.Time
	endPoint   time.Time
	flags      map[int]time.Time
	flagCount  int
}

// NewTimer returns a new Timer with initialized flag map
func NewTimer() *Timer {
	return &Timer{flags: make(map[int]time.Time)}
}

// StartTimer sets startPoint to current time
func (t *Timer) StartTimer() {
	t.startPoint = time.Now()
}

// EndTimer sets endPoint to current time
func (t *Timer) EndTimer() {
	t.endPoint = time.Now()
}

// FlagNow adds currents time to all flags
func (t *Timer) FlagNow() {
	t.flags[t.flagCount] = time.Now()
	t.flagCount++
}

// Show prints the time stamps flags and overall timer time
func (t Timer) Show() {
	for k, v := range t.flags {
		fmt.Println(k, v)
	}
	fmt.Println("\nTotal time: ", t.endPoint.Sub(t.startPoint))
}
