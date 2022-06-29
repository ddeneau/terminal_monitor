package main

import (
	"github.com/rivo/tview"
)

func main() {
	getSystemData()
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}

}

func getSystemData() {

}

/* numberVersions - int - the number of project versions being fed into the test.

 */
