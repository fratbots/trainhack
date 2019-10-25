package main

import (
	"github.com/rivo/tview"
)

const WelcomeMessage = `Ты представитель расы инфузориус вудкюрюус,
твои гениальные собратья эффективно и незаметно
захватили власть над человечеством.

И всё это ради великой цели: создания мегамозга
и перехода в четвёртое измерение.

К сожалению, люди слишком примитивны, чтобы понять
ваш великий замысел, а инфузориус вудкюрюус слишком малы
для достижения такой цели в одиночку.

Поэтому ваш совет старейшин принял судьбоносное решение:
живя в симбиозе с человеком вы вместе должны осуществить
ваш великий замысел.

Не дай своему человеку умереть от болезней и направь
его слабый ум в нужное русло.`

func NewScreenHello() Screen {
	return NewScreen(
		func(game *Game) tview.Primitive {
			return NewUIModal(WelcomeMessage,
				"Play", func() {
					game.SetScreen(NewScreenStage(game, "map2", nil))
				},
				"Exit", func() {
					game.SetScreen(NewScreenFinal())
				},
			)
		},
	)
}
