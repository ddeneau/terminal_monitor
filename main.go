package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

const refreshInterval = 1 * time.Millisecond // Used for application loop
type terminalGUI struct {
	app  *tview.Application
	grid *tview.Grid
}

var (
	memoryTV *tview.TreeView
	cpuTV    *tview.TreeView
)

/* Define structure for each type of system info read-out (modularize CPU, Memory stuff, out of main)*/

func main() {

	ui := initialzeUI()

	memoryTV := tview.NewTreeView() // Don't use treeview here...probably not the best idea.
	memoryTV.SetBackgroundColor(tcell.Color102)
	memoryTV.SetBorder(true).SetTitle("RAM")
	cpuTV := tview.NewTreeView()

	cpuCoreTV := tview.NewTextView()
	cpuTV.SetBackgroundColor(tcell.Color102)
	cpuCoreTV.SetText(getCPU())
	cpuTV.SetBorder(true).SetTitle("CPU")

	ui.grid = tview.NewGrid().AddItem(memoryTV, 0, 0, 1, 1, 1, 1, false)
	ui.grid.AddItem(cpuTV, 1, 0, 1, 1, 1, 1, false)

	if err := ui.app.SetRoot(ui.grid, true).Run(); err != nil {
		panic(err)
	}

	go updateTime(ui.app)

}

/* Start-up the UI with basic components*/
func initialzeUI() terminalGUI {
	ui := terminalGUI{tview.NewApplication(), tview.NewGrid()}
	return ui
}

func updateTime(app *tview.Application) {
	for {
		time.Sleep(refreshInterval)

	}
}

func updateUI() {

}

/* Get the virual memory */
func getVirtualMemory() string {
	v, _ := mem.VirtualMemory()

	total, used, available := v.Total, v.Used, v.Available

	memorySummary := fmt.Sprintf("total: %d mb, \n Used: %d mb, \n Available: %d mb", total, used, available)

	return memorySummary
}

/* Get # cores and percentage used */
func getCPU() string {
	v, _ := cpu.Info()

	coreTime, _ := cpu.Percent(0, false)

	cpuInfo := cpu.InfoStat(v[0])

	cpuSummary := fmt.Sprintf("Cores: %d  Used: %f", cpuInfo.Cores, coreTime[0])

	return cpuSummary
}
