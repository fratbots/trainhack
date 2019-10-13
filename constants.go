package main

import (
	"time"
)

const (
	tickTimeout  = time.Second / 6
	tickTimeoutF = float64(tickTimeout)

	energyGain   = 100
	energyAction = 100
)
