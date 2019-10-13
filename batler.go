package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const noManaImgPath = "./items/nomana.jpg"

type Battle struct {
	Enemy      Character
	Hero       Character
	Log        []string
	Callback   func(screen Screen)
	BackScreen Screen
}

func NewBattle(enemy Character, hero Character, callback func(screen Screen), backScreen Screen) Battle {
	return Battle{
		Enemy:      enemy,
		Hero:       hero,
		Callback:   callback,
		BackScreen: backScreen,
	}
}

type Weapon struct {
	Name    string
	Damage  int
	Energy  int
	ImgPath string
}

type Character interface {
	GetHp() int
	GetMp() int
	GetWeapons() []Weapon
	GetManaRegen() int
	GetImagePath() string
	SetHp(hp int)
	SetMp(mp int)
}

var MagicWand Weapon = Weapon{
	Name:    "Палка Судьбы",
	Damage:  100,
	Energy:  100,
	ImgPath: "./items/posoh.jpg",
}

var DefaultWeapons []Weapon = []Weapon{
	{
		Name:    "Удар Кулаком",
		Damage:  10,
		Energy:  15,
		ImgPath: "./items/kulak.jpg",
	},
	{
		Name:    "Пинок",
		Damage:  15,
		Energy:  25,
		ImgPath: "./items/noga.jpg",
	},
	{
		Name:    "Ждать",
		Damage:  0,
		Energy:  0,
		ImgPath: "./items/zdat.jpg",
	},
}

func (b Battle) Start() tview.Primitive {
	return b.renderBattleScreen(true)
}

func (b Battle) HeroTurn() {
	b.Callback(NewBattleScreen(b.Hero, b.Enemy, b.BackScreen, b.renderBattleScreen(true)))
}

func (b Battle) EnemyTurn() {
	b.reaction(DefaultWeapons[0], false)
}

func (b Battle) DamageScreen(w Weapon, isHeroTurn bool) {
	view := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if isHeroTurn {
				b.Log = append(b.Log, fmt.Sprintf("Вы нанесли врагу %d урона", w.Damage))
				b.EnemyTurn()
			} else {
				b.Log = append(b.Log, fmt.Sprintf("Враг нанес вам %d урона", w.Damage))
				b.HeroTurn()
			}
			return nil
		}), 0, 1, true)
	view.AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(ImageToAscii(w.ImgPath, 120, 40)), 120, 99, true)

	b.Callback(NewBattleScreen(b.Hero, b.Enemy, b.BackScreen, view))
}

func (b Battle) noMana() {
	view := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(ImageToAscii(noManaImgPath, 120, 40)).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			b.HeroTurn()
			return nil
		}), 120, 1, true)

	b.Callback(NewBattleScreen(b.Hero, b.Enemy, b.BackScreen, view))
}

func (b Battle) reaction(w Weapon, isHeroTurn bool) {
	if b.Hero.GetHp() == 0 || b.Enemy.GetHp() == 0 {
		b.Callback(b.BackScreen)
		return
	}

	if isHeroTurn {
		if b.Hero.GetMp() < w.Energy {
			b.noMana()
		}

		b.Enemy.SetHp(b.Enemy.GetHp() - w.Damage)
		b.Hero.SetMp(b.Hero.GetMp() - w.Energy)
	} else {
		b.Hero.SetHp(b.Hero.GetHp() - w.Damage)
	}

	b.DamageScreen(w, isHeroTurn)
}

func (b Battle) renderBattleScreen(isHeroTurn bool) tview.Primitive {
	list := tview.NewList()
	list.AddItem(b.Hero.GetWeapons()[0].Name, fmt.Sprintf("Урон: %d. Энергия: %d", b.Hero.GetWeapons()[0].Damage, b.Hero.GetWeapons()[0].Energy), rune(49), func() { b.reaction(b.Hero.GetWeapons()[0], isHeroTurn) })
	list.AddItem(b.Hero.GetWeapons()[1].Name, fmt.Sprintf("Урон: %d. Энергия: %d", b.Hero.GetWeapons()[1].Damage, b.Hero.GetWeapons()[1].Energy), rune(49), func() { b.reaction(b.Hero.GetWeapons()[1], isHeroTurn) })
	list.AddItem(b.Hero.GetWeapons()[2].Name, fmt.Sprintf("Урон: %d. Энергия: %d", b.Hero.GetWeapons()[2].Damage, b.Hero.GetWeapons()[2].Energy), rune(49), func() { b.reaction(b.Hero.GetWeapons()[2], isHeroTurn) })

	flex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(list, 0, 2, isHeroTurn).
			AddItem(tview.NewTextView().SetText(b.renderBattleLog()).SetTextAlign(tview.AlignCenter), 0, 2, false), 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(fmt.Sprintf("Жизни: %d, Энергия: %d", b.Hero.GetHp(), b.Hero.GetMp())), 0, 5, false).
			AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(ImageToAscii(b.Hero.GetImagePath(), 60, 30)), 0, 45, false).
			AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter), 0, 5, false).
			AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(ImageToAscii(b.Enemy.GetImagePath(), 65, 30)), 0, 45, false).
			AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(fmt.Sprintf("Жизни: %d, Энергия: %d", b.Enemy.GetHp(), b.Enemy.GetMp())), 0, 5, false), 0, 2, false)

	return flex
}

func (b Battle) renderBattleLog() string {
	result := ""
	for _, item := range b.Log {
		result += fmt.Sprintf("%s\n", item)
	}

	return result
}
