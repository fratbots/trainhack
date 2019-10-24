package main

import (
	"github.com/rivo/tview"
)

type BattleScreen struct {
	lastScreen Screen
	enemy      Character
	hero       Character
	primitive  tview.Primitive
}

func (b *BattleScreen) Init(g *Game) tview.Primitive {
	callback := func(s Screen) {
		g.SetScreen(s)
	}
	battle := NewBattle(b.enemy, b.hero, callback, b.lastScreen)
	if b.primitive == nil {
		return battle.Start()
	} else {
		return b.primitive
	}
}

func (b *BattleScreen) Finalize() {}

func NewBattleScreen(hero, enemy Character, lastScreen Screen, primitive tview.Primitive) *BattleScreen {
	return &BattleScreen{
		lastScreen,
		enemy,
		hero,
		primitive,
	}
}
