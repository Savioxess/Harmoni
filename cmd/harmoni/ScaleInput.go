package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var CurrentKeySetTo string = "C"
var CurrentScale string = "Ionian"
var Notes []string = []string{"C", "C#", "D", "D#","E", "F", "F#", "G", "G#", "A", "A#", "B"}
var Scales []string = []string{"Ionian", "Dorian", "Phrygian", "Lydian", "Mixolydian", "Aeolian", "Locrian"}

func createKeyInput() *widget.Select {
  scaleKeyInput := widget.NewSelect(Notes, func(value string) {
    CurrentKeySetTo = value
    fmt.Println("Current Key: ", value)
    RefreshValues()
  })

  return scaleKeyInput
}

func createScaleInput() *widget.Select {
  scaleInput := widget.NewSelect(Scales, func(value string) {
    CurrentScale = value
    fmt.Println("Current Scale: ", value)
    RefreshValues()
  })

  return scaleInput
}

func InputContainer() *fyne.Container {
  keyInput := createKeyInput()
  scaleInput := createScaleInput()
  inputHBox := container.New(layout.NewHBoxLayout(),layout.NewSpacer(), widget.NewLabel("Key"), keyInput, layout.NewSpacer(), widget.NewLabel("Scale"), scaleInput, layout.NewSpacer())

  return inputHBox
}


