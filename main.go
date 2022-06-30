package main

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {

	app := tview.NewApplication()

	total, used, available := getVirtualMemory()
	cpu := getCPU()

	memorySummary := fmt.Sprintf("total: %d, Used: %d, Available: %d", total, used, available)
	cpuSummary := fmt.Sprintf("CPU: %s", cpu)

	memoryText := tview.NewTextView().SetText(memorySummary)
	memoryText.SetTitle("Memory:")
	cpuText := tview.NewTextView().SetText(cpuSummary)
	cpuText.SetTitle("CPU:")


	grid := tview.NewGrid().AddItem(memoryText, 0, 0, 2, 2, 2, 2, false)
	grid.AddItem(cpuText, 1, 0, 2, 2, 2, 2, false)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}

}

func getVirtualMemory() (int64, int64, int64) {
	v, _ := mem.VirtualMemory()

	return int64(v.Total), int64(v.Used), int64(v.Available)
}

func getCPU() string {
	v, _ := cpu.Info()

	cpuInfo := cpu.InfoStat(v[0])

	return cpuInfo.Family
}
