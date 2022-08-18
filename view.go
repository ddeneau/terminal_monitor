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
	app     *tview.Application
	grid    *tview.Grid
	cpuGrid *tview.Grid
	bar     *tview.Grid
	ram     *RAMView
	cpu     *CPUView
	blank   *tview.TextView
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
	ui.cpu, ui.cpuGrid = initializeCPUUI(ui.grid) // Using dual return statements to initialize cpuGrid AND CPUView
	ui.blank = intializeBlankUI(ui.grid)
	ui.bar = tview.NewGrid()
}

func initializeCPUUI(grid *tview.Grid) (*CPUView, *tview.Grid) {
	cpuView := CPUView{}
	cpuGrid := tview.NewGrid()
	percentageGrid := tview.NewGrid()
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
	cpuView.coresTV.SetText("initializing text field").SetTextAlign(1)
	cpuView.coresTV.SetTextColor(tcell.ColorCadetBlue)

	clockTimeLabel.SetText("Clock Time: ")

	cpuView.ctTV.SetText("initializing text field")
	cpuView.ctTV.SetTextColor(tcell.ColorLime)

	percentageGrid.SetBackgroundColor(tcell.Color200)

	modelName.SetText(getCPUTitle())
	modelName.SetTextAlign(1)

	cpuGrid.SetBorder(true)

	percentageGrid.AddItem(cpuView.ctTV, 0, 0, 1, 1, 0, 0, false)
	percentageGrid.AddItem(tview.NewTextView().SetText("%"), 0, 1, 1, 1, 0, 0, false)

	cpuGrid.AddItem(title, 0, 0, 1, 2, 0, 0, false)           // CPU Title row 0, col 0, l = 1, w = 2/
	cpuGrid.AddItem(coresLabel, 1, 0, 1, 1, 0, 0, false)      // CPU Label row 1, col 1, l = 1, w = 1
	cpuGrid.AddItem(cpuView.coresTV, 2, 0, 1, 1, 0, 0, false) // CPU Label row 2, col 1, l = 1, w = 1
	cpuGrid.AddItem(clockTimeLabel, 1, 1, 1, 1, 0, 0, false)  // CPU Label row: 1, col: 0 l = 1 w = 1.  // CPU Data: row: 2, col: 0 l = 1 w = 1.
	cpuGrid.AddItem(percentageGrid, 2, 1, 1, 1, 0, 0, false)
	cpuGrid.AddItem(modelName, 3, 0, 1, 2, 0, 0, false)

	grid.AddItem(cpuGrid, 0, 0, 4, 1, 0, 0, false)
	/* NOTE: A small grid of width '1 unit' can be placed
	// in the main rooted grid, and still have its own number of columns.
	// The CPU box does this by having one inner grid
	// of width '2 sub-units' , and an inner-inner grid of of width '2 sub-sub units'.
	// In other words, each grid gets in own local count of rows and columns to span */
	return &cpuView, cpuGrid
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
	gb := tview.NewTextView()
	gb1 := tview.NewTextView()
	gb2 := tview.NewTextView()

	ramView.usedTV = tview.NewTextView()
	ramView.availableTV = tview.NewTextView()
	ramView.totalTV = tview.NewTextView()

	title.SetText("RAM").SetBackgroundColor(tcell.Color102)
	title.SetTextAlign(1)

	usedLabel.SetText("In Use: ")
	usedLabel.SetBackgroundColor(tcell.ColorWheat)
	availableLabel.SetText("Free: ")
	availableLabel.SetBackgroundColor(tcell.ColorWheat)
	totalLabel.SetText("Total: ")
	totalLabel.SetBackgroundColor(tcell.ColorWheat)

	ramView.usedTV.SetText("initializing text field").SetTextAlign(2)
	ramView.usedTV.SetTextColor(tcell.ColorTomato)
	ramView.usedTV.SetBackgroundColor(tcell.ColorWheat)
	ramView.availableTV.SetText("initializing text field").SetTextAlign(2)
	ramView.availableTV.SetTextColor(tcell.ColorPaleGreen)
	ramView.availableTV.SetBackgroundColor(tcell.ColorWheat)
	ramView.totalTV.SetText("initializing text field").SetTextAlign(2)
	ramView.totalTV.SetTextColor(tcell.ColorPaleGoldenrod)
	ramView.totalTV.SetBackgroundColor(tcell.ColorWheat)

	gb.SetBackgroundColor(tcell.ColorWheat)
	gb.SetText(" GB")

	gb1.SetBackgroundColor(tcell.ColorWheat)
	gb1.SetText(" GB")
	gb2.SetBackgroundColor(tcell.ColorWheat)
	gb2.SetText(" GB")

	memGrid.AddItem(title, 0, 0, 1, 3, 0, 0, false)
	memGrid.AddItem(usedLabel, 1, 0, 1, 1, 0, 0, false)
	memGrid.AddItem(ramView.usedTV, 1, 1, 1, 1, 0, 0, false)
	memGrid.AddItem(gb, 1, 2, 1, 1, 0, 0, false)
	memGrid.AddItem(availableLabel, 2, 0, 1, 1, 0, 0, false)
	memGrid.AddItem(ramView.availableTV, 2, 1, 1, 1, 0, 0, false)
	memGrid.AddItem(gb1, 2, 2, 1, 1, 0, 0, false)
	memGrid.AddItem(totalLabel, 3, 0, 1, 1, 0, 0, false)
	memGrid.AddItem(ramView.totalTV, 3, 1, 1, 1, 0, 0, false)
	memGrid.AddItem(gb2, 3, 2, 1, 1, 0, 0, false)
	grid.AddItem(memGrid, 0, 2, 4, 1, 0, 0, false)
	return &ramView
}

func intializeBlankUI(grid *tview.Grid) *tview.TextView {
	blankGrid := tview.NewGrid()
	title := tview.NewTextView()
	text := tview.NewTextView()
	//table := tview.NewTable()

	title.SetTitle("Processes")
	text.SetText("- in progress -")
	text.SetTextAlign(1)
	blankGrid.SetBorder(true)

	blankGrid.AddItem(text, 0, 0, 1, 1, 1, 1, false)
	grid.AddItem(blankGrid, 4, 0, 2, 3, 0, 0, false)
	return text
}

// Renders system information onto graphical components of a running app.
func updateUI(ui *terminalGUI) {
	updateCpuUI(ui)
	updateRAMUI(ui)
}

func fillBars(amount int8, ui *terminalGUI) {
	color := tcell.NewRGBColor(0, 0, 0)

	for i := 0; i <= int(amount); i++ {
		box := tview.NewBox()

		switch i {
		case 1:
			color = tcell.NewRGBColor(0, 255, 0)
		case 2:
			color = tcell.NewRGBColor(0, 205, 0)
		case 3:
			color = tcell.NewRGBColor(0, 165, 0)
		case 4:
			color = tcell.NewRGBColor(0, 115, 0)

		}

		box.SetBackgroundColor(color)
		ui.bar.AddItem(box, 0, 0, i, 1, 0, 0, false)
	}

	ui.bar.SetBorder(false)
	ui.grid.AddItem(ui.bar, 0, 1, 1, 1, 0, 0, false)
}

func updateCpuUI(ui *terminalGUI) {
	cores, clockTime := getCPU()

	ui.cpu.coresTV.SetText(fmt.Sprintf("%d", cores))
	ui.cpu.ctTV.SetText(fmt.Sprintf("%.4f ", clockTime))

	fillBars(int8(clockTime), ui)

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
