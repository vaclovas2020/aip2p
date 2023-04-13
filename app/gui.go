package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/libp2p/go-libp2p/core/host"
	"webimizer.dev/aip2p/core"
	"webimizer.dev/aip2p/core/node"
)

// Application GUI object
type Gui struct {
	// Main application object
	app fyne.App
	// Main application window
	window fyne.Window
	// Main application text label widget
	text *widget.Label
	// Start button widget
	startBtn *widget.Button
	// libp2p Node
	node host.Host
}

// Start GUI Application main window
func (gui *Gui) Start() {
	gui.app = app.New()
	gui.window = gui.app.NewWindow(fmt.Sprintf("AIP2P Application %s build %d", core.VERSION, core.BUILD_NUMBER))

	gui.text = widget.NewLabel("Start new AIP2P node with \"Start\" button.")
	gui.startBtn = widget.NewButton("Start", func() {
		gui.startBtn.Disable()
		gui.text.SetText("Starting node...")
		peer, err := node.Listen()
		if err != nil {
			panic(err)
		}
		gui.node = peer
		gui.text.SetText(fmt.Sprintf("Listen on %v", peer.Addrs()))
	})
	gui.window.SetContent(container.NewVBox(gui.text, gui.startBtn))
	gui.window.Resize(fyne.NewSize(300, 300))
	gui.window.SetMaster()
	gui.window.ShowAndRun()
	if gui.node == nil {
		return
	}
	if err := gui.node.Close(); err != nil {
		panic(err)
	}
}
