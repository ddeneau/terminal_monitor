package main

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

const refreshInterval = 500 * time.Millisecond // Used for application loop
type terminalGUI struct {
	app  *tview.Application
	grid *tview.Grid
}

var (
	memoryTV *tview.TextView
	cpuTV    *tview.TextView
)

/* Define structure for each type of system info read-out (modularize CPU, Memory stuff, out of main)*/

func main() {

	ui := initialzeUI()

	memoryTV := tview.NewTextView()
	cpuTV := tview.NewTextView()

	ui.grid = tview.NewGrid().AddItem(memoryTV, 0, 0, 2, 2, 2, 2, false)
	ui.grid.AddItem(cpuTV, 1, 0, 2, 2, 2, 2, false)

	if err := ui.app.SetRoot(ui.grid, true).Run(); err != nil {
		panic(err)
	}

	ui.app.Run()
	updateTime(ui.app)

}

/* Start-up the UI with basic components*/
func initialzeUI() terminalGUI {
	ui := terminalGUI{tview.NewApplication(), tview.NewGrid()}
	return ui
}

func updateTime(app *tview.Application) {
	for {
		time.Sleep(refreshInterval)
		app.QueueUpdateDraw(func() {
			updateUI(app)
		})
	}
}

func updateUI(app *tview.Application) {
	memoryTV.SetText(getVirtualMemory())
	cpuTV.SetText(getCPU())
}

/* Get the virual memory */
func getVirtualMemory() string {
	v, _ := mem.VirtualMemory()

	total, used, available := v.Total, v.Used, v.Available

	memorySummary := fmt.Sprintf("total: %d, Used: %d, Available: %d", total, used, available)

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
