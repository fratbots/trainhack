package main

type EffectBehaviorAdd int
type EffectBehaviorFollow int

const (
	// AddBehavior specifies behavior of an effect in case of its repeated initialization.
	EffectBehaviorAddOnce EffectBehaviorAdd = iota
	EffectBehaviorAddRefresh
	EffectBehaviorAddMultiple

	// Target specifies
	EffectBehaviorFollowActor EffectBehaviorFollow = iota
	EffectBehaviorFollowNothing
)

// EffectBehavior represents the ways effect behaves in a specific situations.
type EffectBehavior interface {
	GetAddBehavior() EffectBehaviorAdd
	GetFollowBehavior() EffectBehaviorFollow
}

type EventBehaviorBuilder struct {
	addBehavior    EffectBehaviorAdd
	followBehavior EffectBehaviorFollow
}

func NewEffectBehaviorBuilder(addBehavior EffectBehaviorAdd, followBehavior EffectBehaviorFollow) EventBehaviorBuilder {
	return EventBehaviorBuilder{
		addBehavior:    addBehavior,
		followBehavior: followBehavior,
	}
}

func (b EventBehaviorBuilder) GetAddBehavior() EffectBehaviorAdd {
	return b.addBehavior
}

func (b EventBehaviorBuilder) GetFollowBehavior() EffectBehaviorFollow {
	return b.followBehavior
}
