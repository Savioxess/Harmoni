package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var outputDisplay [8]*fyne.Container
var scale [8]string

func CreateOutputContainer() *fyne.Container {
  createOutputDisplay()
  notesContainer := container.NewHBox(outputDisplay[0], outputDisplay[1], outputDisplay[2], outputDisplay[3], outputDisplay[4], outputDisplay[5], outputDisplay[6], outputDisplay[7])

  notesContainerCentered := container.New(layout.NewCenterLayout(), notesContainer)


  return notesContainerCentered
}

func createOutputDisplay() {
  notesInScale := getNotesInScale()      

  for i, note := range notesInScale {
    outputDisplay[i] = container.NewVBox(widget.NewRichTextFromMarkdown(fmt.Sprintf("## %s", note)), widget.NewLabel(fmt.Sprintf("%d", i+1)))
  }
}

func RefreshValues() {
  notesInScale := getNotesInScale()      

  for i, note := range notesInScale {
    outputDisplay[i].RemoveAll()
    outputDisplay[i].Add(widget.NewRichTextFromMarkdown(fmt.Sprintf("# %s", note)))
    outputDisplay[i].Add(widget.NewLabel(fmt.Sprintf("%d", i+1)))
    outputDisplay[i].Refresh()
  }
}

func getNotesInScale() [8]string {
  var stepsInIonian [7]int = [...]int{2, 2, 1, 2, 2, 2, 1}
  indexOfCurrentScale := getIndexOfCurrentScale()
  indexOfRootNote := getIndexOfRootNode()
  currentNoteIndex := indexOfRootNote

  for i := 0; i < 8; i++ {
    if i-1 != -1 {
      scale[i] = flatOrSharp(scale[i-1], currentNoteIndex)
    } else {
      scale[i] = Notes[currentNoteIndex]
    }
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

func flatOrSharp(previous string, currentNoteIndex int) string {
  noteHalfStepDown := strings.Trim(Notes[currentNoteIndex], "#")
  previousNoteTrim := strings.ReplaceAll(strings.ReplaceAll(previous, "#", ""), "b", "")

  if noteHalfStepDown == previousNoteTrim {
    return Notes[currentNoteIndex+1] + "b" 
  }

  return Notes[currentNoteIndex]
}
