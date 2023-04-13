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

// Start label text
const START_LABEL_TEXT string = "Start new AIP2P node with \"Start\" button."

// Start GUI Application main window
func (gui *Gui) Start() {
	go node.ReceiveExitSignal(gui.node)
	gui.app = app.New()
	gui.window = gui.app.NewWindow(fmt.Sprintf("AIP2P Application %s build %d", core.VERSION, core.BUILD_NUMBER))

	gui.text = widget.NewLabel(START_LABEL_TEXT)
	gui.startBtn = widget.NewButton("Start", func() {
		if gui.node != nil {
			node.StopListen(gui.node)
			gui.node = nil
			gui.startBtn.SetText("Start")
			gui.text.SetText(START_LABEL_TEXT)
			return
		}
		gui.text.SetText("Starting node...")
		peer, err := node.Listen()
		if err != nil {
			panic(err)
		}
		gui.node = peer
		gui.startBtn.SetText("Stop")
		gui.text.SetText(fmt.Sprintf("Listen on %v", peer.Addrs()))
	})
	gui.window.SetContent(container.NewVBox(gui.text, gui.startBtn))
	gui.window.Resize(fyne.NewSize(400, 400))
	gui.window.SetMaster()
	gui.window.ShowAndRun()
	if gui.node != nil {
		node.StopListen(gui.node)
	}
}
