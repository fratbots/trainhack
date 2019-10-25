package main

type Actor struct {
	Class Class

	Position Position
	Energy   Energy

	Behavior    Behavior
	Interaction Interaction

	Hp        int
	Mp        int
	Weapons   []Weapon
	ManaRegen int
	ImagePath string
}

type Behavior func() *Action

type Interaction func(actor *Actor) *Action

func (a *Actor) SetNextAction(action *Action) {
	a.Behavior = func() *Action {
		a.Behavior = nil
		return action
	}
}

func NewHero() *Actor {
	return &Actor{
		Class: Class{
			Name:   "hero",
			IsHero: true,
			Rune:   '@',
			Speed:  1,
		},

		Position: Position{0, 0},
		Energy:   Energy{Value: energyAction}, // full

		Behavior:    nil,
		Interaction: nil,


		// TODO rid off:
		Hp:        100,
		Mp:        100,
		ManaRegen: 5,
		ImagePath: "./example/hero.png",
		Weapons:   DefaultWeapons,
	}
}

func NewActor(pos Position, speed float64, rune rune) *Actor {
	return &Actor{
		Class: Class{
			Name:   "actor",
			IsHero: false,
			Rune:   rune,
			Speed:  speed,
		},

		Position: pos,
		Energy:   Energy{Value: 0}, // empty

		Behavior:    nil,
		Interaction: nil,

		Hp:        120,
		Mp:        20,
		ManaRegen: 1,
		ImagePath: "./example/a.jpeg",
		Weapons:   DefaultWeapons,
	}
}

func (a *Actor) GetHp() int {
	return a.Hp
}

func (a *Actor) GetMp() int {
	return a.Mp
}

func (a *Actor) GetWeapons() []Weapon {
	return a.Weapons
}

func (a *Actor) GetManaRegen() int {
	return a.ManaRegen
}

func (a *Actor) GetImagePath() string {
	return a.ImagePath
}

func (a *Actor) SetHp(hp int) {
	a.Hp = hp
}

func (a *Actor) SetMp(mp int) {
	a.Mp = mp
}
