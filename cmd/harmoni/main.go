package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
  appInstance := app.New()
  window := appInstance.NewWindow("Harmoni")

  window.SetContent(createApplicationContainer())

  window.Show()
  appInstance.Run()
}

func createApplicationContainer() *fyne.Container {
  label := widget.NewLabel("Find Diatonic Chords In Any Mode")
  label.TextStyle.Bold = true
  
  labelContainer := container.New(layout.NewCenterLayout(), label)

  inputContainer := InputContainer() 
  scaleNotesOutputContainer := CreateOutputContainer()
  applicationContainer := container.New(layout.NewVBoxLayout(), labelContainer, inputContainer, scaleNotesOutputContainer)

  return applicationContainer
}
