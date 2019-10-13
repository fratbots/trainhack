package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/naoina/toml"
	"github.com/rivo/tview"
	"io/ioutil"
	"os"
)

type DialogManager struct {
	Dialogs          map[string]Dialog
	CharacterImgPath string
	CharacterName    string
}

type Dialog struct {
	Name   string
	Scenes []Scene
}

type Scene struct {
	Id            int
	PersonName    string
	PersonImgPath string
	Phase         string
	AnswerOne     Answer
	AnswerTwo     Answer
	AnswerThree   Answer
	AnswerFour    Answer
}

type Answer struct {
	Phase   string
	SceneId int
}

func NewDialoger(dialogDirPath string, characterImgPath string, characterName string) DialogManager {
	dialogs := make(map[string]Dialog)

	files, err := ioutil.ReadDir(dialogDirPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		var dialog Dialog
		f, err := os.Open(fmt.Sprintf("%s/%s", dialogDirPath, file.Name()))
		if err != nil {
			panic(err)
		}

		if err := toml.NewDecoder(f).Decode(&dialog); err != nil {
			panic(err)
		}

		dialogs[dialog.Name] = dialog

		err = f.Close()
		if err != nil {
			panic(err)
		}
	}

	return DialogManager{Dialogs: dialogs, CharacterImgPath: characterImgPath, CharacterName: characterName}
}

//func (d DialogManager) StartDialog(dialogName string, render func(s Screen)) tview.Primitive {
//	dialog, ok := d.Dialogs[dialogName]
//	if ok {
//		render(d.getDialogScene(dialog.Scenes, 0))
//	}  else {
//		panic("Диалог не найден")
//	}
//}

//func (d DialogManager) getDialogPrimitive(dialogName string, sceneId int) {
//
//	dialog, ok := d.Dialogs[dialogName]
//	if ok {
//		d.Render(dialog.Scenes, 0)
//	} else {
//		panic("Диалог не найден")
//	}
//
//}

func (d DialogManager) GetDialogPrimitive(dialogName string, sceneId int, callback func(s Screen), lastScreen Screen) tview.Primitive {
	dialog, ok := d.Dialogs[dialogName]
	if !ok {
		panic("Диалог не найден")
	}
	scene := dialog.Scenes[sceneId]
	characterImgAscii := ImageToAscii(d.CharacterImgPath, 65, 23)
	personImgAscii := ImageToAscii(scene.PersonImgPath, 65, 23)

	//newPrimitive := func(text string) tview.Primitive {
	//	return tview.NewTextView().
	//		SetTextAlign(tview.AlignCenter).
	//		SetText(text)
	//}

	list := tview.NewList()
	list.AddItem(scene.AnswerOne.Phase, "", '1', func() { newScene(scene.AnswerOne.SceneId, dialogName, lastScreen, callback) })
	list.AddItem(scene.AnswerTwo.Phase, "", '2', func() { newScene(scene.AnswerTwo.SceneId, dialogName, lastScreen, callback) })
	list.AddItem(scene.AnswerThree.Phase, "", '3', func() { newScene(scene.AnswerThree.SceneId, dialogName, lastScreen, callback) })
	list.AddItem(scene.AnswerFour.Phase, "", '4', func() { newScene(scene.AnswerFour.SceneId, dialogName, lastScreen, callback) })

	list.SetShortcutColor(tcell.Color18)
	list.SetBackgroundColor(tcell.ColorDimGray)

	//grid := tview.NewGrid().
	//	SetRows(20, 1, 7).
	//	SetColumns(10, 0, 10).
	//	SetBorders(true)
	//
	//grid.AddItem(newPrimitive(fmt.Sprintf(characterImgAscii)), 0, 0, 1, 2, 0, 0, false)
	//grid.AddItem(newPrimitive(personImgAscii), 0, 2, 1, 2, 0, 0, false)
	//
	//grid.AddItem(newPrimitive(d.CharacterName), 1, 0, 1, 2, 0, 0, false)
	//grid.AddItem(newPrimitive(scene.PersonName), 1, 2, 1, 2, 0, 0, false)
	//
	//grid.AddItem(list, 2, 0, 1, 2, 0, 0, true)
	//grid.AddItem(newPrimitive(scene.Phase), 2, 2, 1, 2, 0, 0, false)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(tview.NewTextView().SetText(characterImgAscii), 65, 4, false).
			AddItem(tview.NewBox(), 0, 2, false).
			AddItem(tview.NewTextView().SetText(personImgAscii), 65, 4, false),
			23, 6, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(tview.NewTextView().SetText(d.CharacterName).SetTextAlign(tview.AlignCenter), 0, 1, false).
			AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("|"), 0, 1, false).
			AddItem(tview.NewTextView().SetText(scene.PersonName).SetTextAlign(tview.AlignCenter), 0, 1, false), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(list, 0, 4, true).
			AddItem(tview.NewTextView().SetText("|").SetTextAlign(tview.AlignCenter), 0, 2, false).
			AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(scene.Phase), 0, 4, false), 0, 3, true)

	return flex
}

func newScene(sceneId int, dialogName string, lastScreen Screen, callback func(s Screen)) {
	if sceneId == -1 {
		callback(lastScreen)
		return
	}
	callback(NewDialogScreen(dialogName, sceneId, lastScreen))
}
