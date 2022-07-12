package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type terminalGUI struct {
	app       *tview.Application
	grid      *tview.Grid
	memText   *tview.TextView
	cpuText   *tview.TextView
	tempTimer *tview.TextView
}

var refreshRate = 1 * time.Second
var globalTimer = 0

func initializeApp() {
	ui := initialzeUI()

	// Set up quit key input capture.
	ui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			ui.app.Stop()
		}

		return event
	})

	go refresh(&ui) // goroutine for refreshing.

	if err := ui.app.SetRoot(ui.grid, true).Run(); err != nil {
		panic(err)
	}
}

/* Start-up the UI with basic components (two basic components for now, lol)*/
func initialzeUI() terminalGUI {
	app, grid := tview.NewApplication(), tview.NewGrid()
	memText, cpuText, tt := tview.NewTextView(), tview.NewTextView(), tview.NewTextView()

	memText.SetBackgroundColor(tcell.Color102)
	memText.SetBorder(true).SetTitle("RAM")

	cpuText.SetBackgroundColor(tcell.Color102)
	cpuText.SetBorder(true).SetTitle("CPU")

	ui := terminalGUI{app, grid, memText, cpuText, tt}

	return ui
}

// Renders system information onto graphical components of a running app.
func updateUI(ui *terminalGUI) {
	sysInfo := [2]string{getCPU(), getVirtualMemory()}
	ui.cpuText.SetText(sysInfo[0])
	ui.memText.SetText(sysInfo[1])
	ui.tempTimer.SetText(fmt.Sprintf("Terminal Monitor Running For: %d Seconds", globalTimer))
	ui.grid.AddItem(ui.cpuText, 0, 0, 2, 1, 1, 1, false)
	ui.grid.AddItem(ui.memText, 0, 1, 2, 1, 1, 1, false)
	ui.grid.AddItem(ui.tempTimer, 2, 0, 1, 2, 1, 2, false)
	globalTimer += 1
}

// Uses an infinite loop to wait a period of time, then send an update to the running app.
func refresh(ui *terminalGUI) {
	for {
		time.Sleep(refreshRate)
		ui.app.QueueUpdateDraw(func() {
			updateUI(ui)
		})
	}
}
