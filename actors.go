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

	MustBeDeleted bool
	Direction     Direction
}

type Behavior func() *Action

type Interaction func(target *Actor) *Action

func (a *Actor) SetNextAction(action *Action) {
	a.Behavior = func() *Action {
		a.Behavior = nil
		return action
	}
}

func NewClassActor(stage *Stage, pos Position, direction Direction, class string) *Actor {
	classes := Classes()
	cls, ok := classes[class]
	if !ok {
		return nil
	}

	actor := &Actor{
		Class: cls,

		Position: pos,
		Energy:   Energy{Value: 0},

		Behavior:    nil,
		Interaction: nil,

		// нужно для игры с мыслями
		Direction: direction,
	}

	// default behavior
	if actor.Class.BehaviorInit != nil {
		actor.Behavior = actor.Class.BehaviorInit(stage, actor)
	}

	// default interaction
	if actor.Class.InteractionInit != nil {
		actor.Interaction = actor.Class.InteractionInit(stage, actor)
	}

	return actor
}

func GetActorState(actor *Actor) ActorState {
	return ActorState{
		ClassName: actor.Class.Name,
		Position:  actor.Position,
		Energy:    actor.Energy,
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
