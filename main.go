package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)
type terminalGUI struct {
	app        *tview.Application
	grid       *tview.Grid
	memoryView *tview.Box
	cpuView    *tview.Box
	memText    *tview.TextView
	cpuText    *tview.TextView
}


/* Define structure for each type of system info read-out (modularize CPU, Memory stuff, out of main)*/

func main() {

	ui := initialzeUI()

	if err := ui.app.SetRoot(ui.grid, true).Run(); err != nil {
		panic(err)
	}

	updateUI(&ui)

}

/* Start-up the UI with basic components (two basic components for now, lol)*/
func initialzeUI() terminalGUI {
	app, grid := tview.NewApplication(), tview.NewGrid()
	memoryView, cpuView := tview.NewBox(), tview.NewBox()
	memText, cpuText := tview.NewTextView(), tview.NewTextView()

	memoryView.SetBackgroundColor(tcell.Color102)
	memoryView.SetBorder(true).SetTitle("RAM")

	cpuView.SetBackgroundColor(tcell.Color102)
	cpuView.SetBorder(true).SetTitle("CPU")

	ui := terminalGUI{app, grid, memoryView, cpuView, memText, cpuText}

	return ui
}

func updateUI(ui *terminalGUI) {
	sysInfo := [2]string {getCPU(), getVirtualMemory()}
	ui.cpuText.SetText(sysInfo[0])
	ui.memText.SetText(sysInfo[1])
}
