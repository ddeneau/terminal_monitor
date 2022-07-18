package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

/* Contains important global parts of the app, and individual
   sections that get updated. */
type terminalGUI struct {
	app   *tview.Application
	grid  *tview.Grid
	cpu   *tview.TextView
	ram   *tview.TextView
	blank *tview.TextView
}

var refreshRate = 1 * time.Second // Refresh the screen this frequently.

/* Set-up major components, quit button, and start refreshing.  */
func initializeApp() {
	ui := initialzeUI()
	initializeUIComponents(&ui)

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
	app, grid := tview.NewApplication(), tview.NewGrid() // Outermost grid in this grid of grids

	ui := terminalGUI{}
	ui.app = app
	ui.grid = grid

	return ui
}

/* Just initializes each sub-component of the UI. */
func initializeUIComponents(ui *terminalGUI) {
	ui.ram = initializeRAMUI(ui.grid)
	ui.cpu = initializeCPUUI(ui.grid)
	// ui.blank = intializeBlankUI(ui.grid)
}

/* Each of these  just add UI components to the global grid, passed in, and return the
   UI component that needs to be refreshed  */

func initializeRAMUI(grid *tview.Grid) *tview.TextView {
	memGrid := tview.NewGrid()
	title := tview.NewTextView()
	text := tview.NewTextView()

	title.SetBackgroundColor(tcell.Color102)
	title.SetText("RAM")

	text.SetBackgroundColor(tcell.Color102)

	memGrid.SetBorder(true)

	memGrid.AddItem(title, 0, 0, 1, 1, 1, 1, false)
	memGrid.AddItem(text, 1, 0, 1, 1, 1, 1, false)
	grid.AddItem(memGrid, 0, 2, 1, 2, 2, 2, false)

	return text
}

func initializeCPUUI(grid *tview.Grid) *tview.TextView {
	cpuGrid := tview.NewGrid()
	title := tview.NewTextView()
	text := tview.NewTextView()

	title.SetBackgroundColor(tcell.Color102)
	title.SetText(getCPUTitle())
	text.SetBackgroundColor(tcell.Color102)

	cpuGrid.SetBorder(true)

	cpuGrid.AddItem(title, 0, 0, 1, 1, 1, 1, false)
	cpuGrid.AddItem(text, 1, 0, 1, 1, 1, 1, false)
	grid.AddItem(cpuGrid, 0, 0, 1, 2, 2, 2, false)

	return text
}

func intializeBlankUI(grid *tview.Grid) *tview.TextView {
	blankGrid := tview.NewGrid()
	title := tview.NewTextView()
	text := tview.NewTextView()

	blankGrid.AddItem(title, 0, 0, 1, 1, 1, 1, false)
	blankGrid.AddItem(text, 0, 1, 1, 1, 1, 1, false)
	grid.AddItem(blankGrid, 0, 4, 1, 1, 1, 2, false)

	return text
}

// Renders system information onto graphical components of a running app.
func updateUI(ui *terminalGUI) {
	updateCpuUI(ui)
	updateRAMUI(ui)
}

func updateCpuUI(ui *terminalGUI) {
	ui.cpu.SetText(getCPU())
}

func updateRAMUI(ui *terminalGUI) {
	ui.ram.SetText(getVirtualMemory())
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
