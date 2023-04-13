package aip2p

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"webimizer.dev/aip2p/core"
	"webimizer.dev/aip2p/core/node"
)

// AIP2P application entry function
func StartApplication() {
	fmt.Printf("AIP2P Application %s build %d\n\n", core.VERSION, core.BUILD_NUMBER)
	node.StartNode()
}

func StartApplicationGUI() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.Show()

	w2 := a.NewWindow("Larger")
	w2.SetContent(widget.NewLabel("More content"))
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()

	a.Run()
}
