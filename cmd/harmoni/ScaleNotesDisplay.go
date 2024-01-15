package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var outputDisplay *widget.Label

func CreateOutputContainer() *fyne.Container {
  createOutputDisplay()
  outputContainer := container.New(layout.NewCenterLayout(), outputDisplay)

  return outputContainer
}

func createOutputDisplay() {
  notesInScale := getNotesInScale()      
  outputText := fmt.Sprintf("%s %s %s %s %s %s %s %s", notesInScale[0], notesInScale[1], notesInScale[2], notesInScale[3], notesInScale[4], notesInScale[5], notesInScale[6], notesInScale[7]) 

  outputDisplay = widget.NewLabel(outputText)
}

func RefreshValues() {
  notesInScale := getNotesInScale()      
  outputText := fmt.Sprintf("%s %s %s %s %s %s %s %s", notesInScale[0], notesInScale[1], notesInScale[2], notesInScale[3], notesInScale[4], notesInScale[5], notesInScale[6], notesInScale[7]) 

  outputDisplay.SetText(outputText)
  outputDisplay.Refresh()
}

func getNotesInScale() [8]string {
  var scale [8]string
  var stepsInIonian [7]int = [...]int{2, 2, 1, 2, 2, 2, 1}
  indexOfCurrentScale := getIndexOfCurrentScale()
  indexOfRootNote := getIndexOfRootNode()
  currentNoteIndex := indexOfRootNote

  for i := 0; i < 8; i++ {
    scale[i] = Notes[currentNoteIndex]
    currentNoteIndex = getNextNodeIndex(currentNoteIndex + stepsInIonian[getStepIndex(i + indexOfCurrentScale)])
  }

  return scale
}

func getStepIndex(index int) int {
  if index >= 7 {
    return index - 7
  }

  return index
}

func getNextNodeIndex(index int) int {
  if index >= 12 {
    return index - 12
  }

  return index
}

func getIndexOfCurrentScale() int {
  for index, value := range Scales {
    if value == CurrentScale {
      return index
    }
  }

  return 0
}

func getIndexOfRootNode() int {
  for index, value := range Notes {
    if value == CurrentKeySetTo {
      return index
    }
  }

  return 0
}
