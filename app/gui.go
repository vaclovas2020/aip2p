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
	// Copy button widget
	copyBtn *widget.Button
	// libp2p Node
	node *host.Host
	// log text string
	logText string
	// current node Address
	address string
}

// Start label text
const START_LABEL_TEXT string = "Start new AIP2P node with \"Start\" button."

// Start GUI Application main window
func (gui *Gui) Start() {
	_, err := os.Open("aip2p.lock")
	if err != nil && os.IsNotExist(err) {
		f, err := os.Create("aip2p.lock")
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
				gui.copyBtn.Disable()
				return
			}
			peer, err := node.Listen()
			if err != nil {
				panic(err)
			}
			addrs, err := node.GetAddress(peer)
			if err != nil {
				panic(err)
			}
			gui.node = peer
			gui.address = addrs[0].String()
			gui.startBtn.SetText("Stop")
			gui.logText += fmt.Sprintf("libp2p node address: %v\n", addrs[0])
			gui.text.SetText(gui.logText)
			gui.copyBtn.Enable()
		})
		gui.copyBtn = widget.NewButton("Copy Address", func() {
			if gui.address != "" {
				gui.window.Clipboard().SetContent(gui.address)
			}
		})
		gui.copyBtn.Disable()
		gui.window.SetContent(container.NewVBox(gui.text, gui.startBtn, gui.copyBtn))
		gui.window.Resize(fyne.NewSize(600, 600))
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
		os.Remove("aip2p.lock")
	}
}
