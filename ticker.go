package main

import (
	"time"
)

type Ticker struct {
	ticker *time.Ticker
	last   time.Time
	ch     chan bool
}

func NewTicker(timeout time.Duration, callback func(delta time.Duration)) *Ticker {
	ticker := time.NewTicker(timeout)

	t := &Ticker{
		ticker: ticker,
		last:   time.Now(),
		ch:     make(chan bool),
	}

	go func() {
		for {
			select {
			case b := <-t.ch:
				if b == false {
					return
				}
			case <-ticker.C:
			}

			callback(time.Since(t.last))
			t.last = time.Now()
		}
	}()

	return t
}

func (t *Ticker) Tick() {
	t.ch <- true
}

func (t *Ticker) Done() {
	t.ch <- false
}
