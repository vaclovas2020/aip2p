package app

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/libp2p/go-libp2p/core/host"
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
	node *host.Host
	// log text string
	logText string
}

// Start label text
const START_LABEL_TEXT string = "Start new AIP2P node with \"Start\" button."

// Start GUI Application main window
func (gui *Gui) Start() {
	_, err := os.Open(".lock")
	if err != nil && os.IsNotExist(err) {
		f, err := os.Create(".lock")
		if err != nil {
			f.Write([]byte("1"))
		}
		gui.app = app.New()
		gui.window = gui.app.NewWindow(fmt.Sprintf("AIP2P Application v%s build %d", gui.app.Metadata().Version, gui.app.Metadata().Build))

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
			gui.logText += fmt.Sprintf("Listen on %v\n", (*peer).Addrs()[0])
			gui.text.SetText(gui.logText)
		})
		gui.window.SetContent(container.NewVBox(gui.text, gui.startBtn))
		gui.window.Resize(fyne.NewSize(600, 400))
		gui.window.SetFixedSize(true)
		gui.window.CenterOnScreen()
		gui.window.SetMaster()
		if desk, ok := gui.app.(desktop.App); ok {
			m := fyne.NewMenu("AIP2P",
				fyne.NewMenuItem("Show", func() {
					gui.window.Show()
				}))
			desk.SetSystemTrayMenu(m)
			gui.window.SetCloseIntercept(func() {
				gui.window.Hide()
			})
		}
		if len(os.Args) > 1 && os.Args[1] == "systray" {
			gui.app.Run()
		} else {
			gui.window.ShowAndRun()
		}
		if gui.node != nil {
			node.StopListen(gui.node)
		}
		os.Remove(".lock")
	}
}
