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

func (b *BattleScreen) Do(g *Game, callback func(s Screen)) tview.Primitive {
	battle := NewBattle(b.enemy, b.hero, callback, b.lastScreen)
	if b.primitive == nil {
		return battle.Start()
	} else {
		return b.primitive
	}
}

func NewBattleScreen(hero, enemy Character, lastScreen Screen, primitive tview.Primitive) *BattleScreen {
	return &BattleScreen{
		lastScreen,
		enemy,
		hero,
		primitive,
	}
}
