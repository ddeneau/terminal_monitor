package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

/* Contains important global parts of the app, and individual
   sections that get updated. */

type CPUView struct {
	coresTV *tview.TextView
	ctTV    *tview.TextView
}

type RAMView struct {
	totalTV     *tview.TextView
	usedTV      *tview.TextView
	availableTV *tview.TextView
}
type terminalGUI struct {
	app   *tview.Application
	grid  *tview.Grid
	ram   *RAMView
	cpu   *CPUView
	blank *tview.TextView
}

var refreshRate = 1 * time.Second // Refresh the screen this frequently.

/* Set-up major components, quit button, and start refreshing.
1) ui global object gets created and set to the return variable from the initialization function
2) app and grid components are set up and set to the ui global object, which is passed to the global object.
3) components are set up in sub-routine, which takes the ui address as a parameter.
4) global ui object gets assigned major values from function return values which contain themselves objects per UI section
5a) - initRAMUI: returns the RAMView object. Sets up static text fields.
5b) - initCPUUI:  returns the CPUView object. Sets up static text fields.
6) Main event loop starts:
7) Update is called as a go routine, with a one second thread sleep timer.
8) - values of global UI sub-structures are updated (TextViews need to be updated.)


*/
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
	ui.blank = intializeBlankUI(ui.grid)
}

/* Each of these  just add UI components to the global grid, passed in, and return the
   UI component that needs to be refreshed  */

func initializeRAMUI(grid *tview.Grid) *RAMView {
	ramView := RAMView{}
	memGrid := tview.NewGrid()
	title := tview.NewTextView()
	usedLabel := tview.NewTextView()
	availableLabel := tview.NewTextView()
	totalLabel := tview.NewTextView()

	ramView.usedTV = tview.NewTextView()
	ramView.availableTV = tview.NewTextView()
	ramView.totalTV = tview.NewTextView()

	title.SetText("RAM").SetBackgroundColor(tcell.Color102)
	title.SetTextAlign(1)

	usedLabel.SetText("In Use: ")
	availableLabel.SetText("Free: ")
	totalLabel.SetText("Total: ")

	ramView.usedTV.SetText("initializing text field").SetTextAlign(2).SetTextColor(tcell.ColorTomato)
	ramView.availableTV.SetText("initializing text field").SetTextAlign(2).SetTextColor(tcell.ColorPaleGreen)
	ramView.totalTV.SetText("initializing text field").SetTextAlign(2).SetTextColor(tcell.ColorPaleGoldenrod)

	memGrid.SetBorder(true)

	memGrid.AddItem(title, 0, 0, 1, 3, 0, 0, false)
	memGrid.AddItem(usedLabel, 1, 0, 1, 1, 0, 0, false)
	memGrid.AddItem(ramView.usedTV, 1, 1, 1, 1, 0, 0, false)
	memGrid.AddItem(tview.NewTextView().SetText("GB").SetTextAlign(1), 1, 2, 1, 1, 0, 0, false)
	memGrid.AddItem(availableLabel, 2, 0, 1, 1, 0, 0, false)
	memGrid.AddItem(ramView.availableTV, 2, 1, 1, 1, 0, 0, false)
	memGrid.AddItem(tview.NewTextView().SetText("GB").SetTextAlign(1), 2, 2, 1, 1, 0, 0, false)
	memGrid.AddItem(totalLabel, 3, 0, 1, 1, 0, 0, false)
	memGrid.AddItem(ramView.totalTV, 3, 1, 1, 1, 0, 0, false)
	memGrid.AddItem(tview.NewTextView().SetText("GB").SetTextAlign(1), 3, 2, 1, 1, 0, 0, false)
	grid.AddItem(memGrid, 0, 2, 4, 2, 0, 0, false)
	return &ramView
}

func initializeCPUUI(grid *tview.Grid) *CPUView {
	cpuView := CPUView{}

	cpuGrid := tview.NewGrid()
	title := tview.NewTextView()
	clockTimeLabel := tview.NewTextView()
	coresLabel := tview.NewTextView()
	modelName := tview.NewTextView()
	cpuView.coresTV = tview.NewTextView()
	cpuView.ctTV = tview.NewTextView()

	//title.SetBackgroundColor(tcell.Color102)
	title.SetText("CPU").SetBackgroundColor(tcell.Color102)
	title.SetTextAlign(1)

	coresLabel.SetText("Cores: ").SetTextAlign(1)
	cpuView.coresTV.SetText("initializing text field").SetTextAlign(1).SetTextColor(tcell.ColorCadetBlue)

	clockTimeLabel.SetText("Clock Time: ").SetTextAlign(2)
	cpuView.ctTV.SetText("initializing text field").SetTextColor(tcell.ColorLime).SetTextAlign(2)

	modelName.SetText(getCPUTitle()).SetTitleAlign(1)

	cpuGrid.SetBorder(true)

	cpuGrid.AddItem(title, 0, 0, 1, 3, 0, 0, false)           // CPU Title row 0, col 0, l = 1, w = 2/
	cpuGrid.AddItem(coresLabel, 1, 0, 1, 1, 0, 0, false)      // CPU Label row 1, col 1, l = 1, w = 1
	cpuGrid.AddItem(cpuView.coresTV, 2, 0, 1, 1, 0, 0, false) // CPU Label row 2, col 1, l = 1, w = 1
	cpuGrid.AddItem(clockTimeLabel, 1, 1, 1, 1, 0, 0, false)  // CPU Label row: 1, col: 0 l = 1 w = 1.
	cpuGrid.AddItem(cpuView.ctTV, 2, 1, 1, 1, 0, 0, false)    // CPU Data: row: 2, col: 0 l = 1 w = 1.
	cpuGrid.AddItem(tview.NewTextView().SetText("%"), 2, 2, 1, 1, 0, 0, false)
	cpuGrid.AddItem(modelName, 3, 0, 1, 3, 0, 0, false)

	grid.AddItem(cpuGrid, 0, 0, 4, 2, 0, 0, false)
	return &cpuView
}

func intializeBlankUI(grid *tview.Grid) *tview.TextView {
	blankGrid := tview.NewGrid()
	title := tview.NewTextView()
	text := tview.NewTextView()
	//table := tview.NewTable()

	title.SetTitle("Processes")
	text.SetText("Disk(s) To be implemented")
	text.SetTextAlign(1)
	blankGrid.SetBorder(true)

	blankGrid.AddItem(text, 0, 0, 1, 1, 1, 1, false)
	grid.AddItem(blankGrid, 5, 0, 4, 4, 4, 4, false)
	return text
}

// Renders system information onto graphical components of a running app.
func updateUI(ui *terminalGUI) {
	updateCpuUI(ui)
	updateRAMUI(ui)
}

func updateCpuUI(ui *terminalGUI) {
	cores, clockTime := getCPU()

	ui.cpu.coresTV.SetText(fmt.Sprintf("%d", cores))
	ui.cpu.ctTV.SetText(fmt.Sprintf("%.4f ", clockTime))
	//ui.cpu = &cpuView
}

func updateRAMUI(ui *terminalGUI) {
	ram := getVirtualMemory() // New RAM system call.

	ui.ram.usedTV.SetText(fmt.Sprintf("%.4f", ram.used))
	ui.ram.availableTV.SetText(fmt.Sprintf("%.4f", ram.available))
	ui.ram.totalTV.SetText(fmt.Sprintf("%.4f", ram.total))
	//ui.ram = &ramView
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
