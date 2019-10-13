package main

import (
	"github.com/rivo/tview"
)

const WelcomeMessage = `Ты представитель рассы инфузориус вудкюрюус,
твои гениальные собратья эффективно и незаметно
захватили власть над человечевством.

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

type HelloScreen struct {
}

func (s *HelloScreen) Do(g *Game, end func(next Screen)) tview.Primitive {

	var modal *tview.Modal
	modal = tview.NewModal().
		SetText(WelcomeMessage).
		AddButtons([]string{"Play", "Exit"}).
		SetDoneFunc(
			func(buttonIndex int, buttonLabel string) {
				if buttonIndex == 0 {
					end(NewScreenStage(g, "map2", nil))
					return
				}
				if buttonIndex == 1 {
					end(&ScreenFinal{})
					return
				}

				modal.SetText("You win!")
			},
		)

	return modal
}
