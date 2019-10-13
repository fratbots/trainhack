package main

type Port struct {
	MapOnScreen Position
	Screen      Dimensions
	ScreenShift Dimensions
}

func NewPort(screen, screenShift Dimensions, lookAtMap Position) Port {
	return Port{
		MapOnScreen: Position{
			X: lookAtMap.X - screen.X/2,
			Y: lookAtMap.Y - screen.Y/2,
		},
		Screen:      screen,
		ScreenShift: screenShift,
	}
}

func (p *Port) ToMap(screen Position) (mapPos Position) {
	mapPos.X = p.MapOnScreen.X + screen.X
	mapPos.Y = p.MapOnScreen.Y + screen.Y
	return
}

func (p *Port) ToScreen(mapPos Position) (screen Position) {
	screen.X = mapPos.X - p.MapOnScreen.X + p.ScreenShift.X
	screen.Y = mapPos.Y - p.MapOnScreen.Y + p.ScreenShift.Y
	return
}
