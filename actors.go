package main

type Actor struct {
	IsHero   bool
	Rune     rune
	Position Position
	Energy   Energy
	Speed    float64

	Hp        int
	Mp        int
	Weapons   []Weapon
	ManaRegen int
	ImagePath string

	Behavior    Behavior
	Interaction Interaction
}

type Behavior func() *Action

type Interaction func(actor *Actor) *Action

func (h *Actor) NextAction(action *Action) {
	h.Behavior = func() *Action {
		h.Behavior = nil
		return action
	}
}

func NewHero(weapons []Weapon) *Actor {
	return &Actor{
		IsHero:    true,
		Rune:      '@',
		Position:  Position{0, 0},
		Energy:    Energy{Value: energyAction},
		Speed:     1,
		Behavior:  nil,
		Hp:        100,
		Mp:        100,
		ManaRegen: 5,
		ImagePath: "./example/k.png",
		Weapons:   weapons,
	}
}

func NewActor(pos Position, speed float64, rune rune) *Actor {
	return &Actor{
		IsHero:    false,
		Rune:      rune,
		Position:  pos,
		Energy:    Energy{Value: 0},
		Speed:     speed,
		Behavior:  nil,
		Hp:        20,
		Mp:        0,
		ManaRegen: 0,
		ImagePath: "./example/a.jpeg",
		Weapons:   DefaultWeapons,
	}
}

func (a Actor) GetHp() int {
	return a.Hp
}

func (a Actor) GetMp() int {
	return a.Mp
}

func (a Actor) GetWeapons() []Weapon {
	return a.Weapons
}

func (a Actor) GetManaRegen() int {
	return a.ManaRegen
}

func (a Actor) GetImagePath() string {
	return a.ImagePath
}

func (a Actor) SetHp(hp int) {
	a.Hp = hp
}

func (a Actor) SetMp(mp int) {
	a.Mp = mp
}
