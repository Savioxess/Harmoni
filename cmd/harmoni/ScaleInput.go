package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var currentKeySetTo string = "C"
var currentScale string = "Ionian"

func createKeyInput() *widget.Select {
  var notes []string = []string{"C", "C#", "D", "D#","E", "F", "F#", "G", "G#", "A", "A#", "B"}

  scaleKeyInput := widget.NewSelect(notes, func(value string) {
    currentKeySetTo = value
    fmt.Println("Current Key: ", value)
  })

  return scaleKeyInput
}

func createScaleInput() *widget.Select {
  var scales []string = []string{"Ionian", "Dorian", "Phrygian", "Lydian", "Mixolydian", "Aeolian", "Locrian"}

  scaleInput := widget.NewSelect(scales, func(value string) {
    currentScale = value
    fmt.Println("Current Scale: ", value)
  })

  return scaleInput
}

func InputContainer() *fyne.Container {
  inputHBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), createKeyInput(), layout.NewSpacer(), createScaleInput(), layout.NewSpacer())

  return inputHBox
}


