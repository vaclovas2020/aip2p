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
	text *widget.Entry
	// Start button widget
	startBtn *widget.Button
	// libp2p Node
	node host.Host
	//
	logText string
}

// Start label text
const START_LABEL_TEXT string = "Start new AIP2P node with \"Start\" button."

// Start GUI Application main window
func (gui *Gui) Start() {
	go node.ReceiveExitSignal(gui.node)
	gui.app = app.New()
	gui.window = gui.app.NewWindow(fmt.Sprintf("AIP2P Application %s build %d", core.VERSION, core.BUILD_NUMBER))

	gui.logText = fmt.Sprintln(START_LABEL_TEXT)
	gui.text = widget.NewMultiLineEntry()
	gui.text.SetText(gui.logText)
	gui.text.Disable()
	gui.text.SetMinRowsVisible(10)
	gui.startBtn = widget.NewButton("Start", func() {
		if gui.node != nil {
			node.StopListen(gui.node)
			gui.node = nil
			gui.startBtn.SetText("Start")
			gui.logText += fmt.Sprintln("Stopped.")
			gui.logText += fmt.Sprintln(START_LABEL_TEXT)
			gui.text.SetText(gui.logText)
			return
		}
		peer, err := node.Listen()
		if err != nil {
			panic(err)
		}
		gui.node = peer
		gui.startBtn.SetText("Stop")
		gui.logText += fmt.Sprintf("Listen on %v\n", peer.Addrs())
		gui.text.SetText(gui.logText)
	})
	gui.window.SetContent(container.NewVBox(gui.text, gui.startBtn))
	gui.window.Resize(fyne.NewSize(600, 400))
	gui.window.SetMaster()
	gui.window.ShowAndRun()
	if gui.node != nil {
		node.StopListen(gui.node)
	}
}
