package main

import (
	"time"
)

const (
	tickTimeout      = time.Second / 6
	tickTimeoutFloat = float64(tickTimeout)

	energyGain   = 100
	energyAction = 100
)
