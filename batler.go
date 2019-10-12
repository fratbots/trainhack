package main

import (
	"fmt"
	"github.com/rivo/tview"
)

type Battle struct {
	Enemy      Character
	Hero       Character
	Log        []string
	Callback   func(screen Screen)
	BackScreen Screen
}

func NewBattle(enemy Character, hero Character, callback func(screen Screen), backScreen Screen) {

}

type Weapon struct {
	Name   string
	Damage int
	Energy int
}

type Character interface {
	GetHp() int
	GetMp() int
	GetHandDamage() int
	GetFootDamage() int
	GetWeapons() []Weapon
	GetManaRegen() int
	GetImagePath() string
	SetHp(hp int)
	SetMp(mp int)
}

func (b Battle) BattleLoop() {
	for {
		if b.Hero.GetHp() == 0 || b.Enemy.GetHp() == 0 {
			b.Callback(b.BackScreen)
			return
		}

		b.Hero.SetMp(b.Hero.GetMp() + b.Hero.GetManaRegen())
	}

}

func (b Battle) HeroTurn() tview.Primitive {
	list := tview.NewList()
	list.AddItem("Удар Рукой", fmt.Sprintf("Урон: %d. Энергия: %d", b.Hero.GetHandDamage(), 10), '1', func() {})
	list.AddItem("Удар ногой", fmt.Sprintf("Урон: %d. Энергия: %d", b.Hero.GetFootDamage(), 20), '2', func() {})
	list.AddItem("Ничего не делать", "", '3', func() {})
	for i, item := range b.Hero.GetWeapons() {
		list.AddItem(item.Name, fmt.Sprintf("Урон: %d. Энергия: %d", item.Damage, item.Energy), rune(i+4), func() {})
	}

}

func (b Battle) EnemyTurn() tview.Primitive {

}

func (b Battle) DamageScreen() tview.Primitive {

}

func NoManaPrimitive() tview.Primitive {

}

func Reaction() {

}
