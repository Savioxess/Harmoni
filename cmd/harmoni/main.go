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
  label := widget.NewRichTextFromMarkdown("# Chords Finder")
  
  labelContainer := container.NewCenter(label)

  inputContainer := InputContainer() 
  scaleNotesOutputContainer := CreateOutputContainer()
  applicationContainer := container.New(layout.NewVBoxLayout(), labelContainer, layout.NewSpacer(), inputContainer, scaleNotesOutputContainer, layout.NewSpacer())

  return applicationContainer
}
