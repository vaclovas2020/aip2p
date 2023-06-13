// Implements AIP2P GUI Application based on fyne.io/fyne
package app

import (
	"fmt"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
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
	// Connections list widget
	connectionsTable *widget.Table
	// Connect text widget
	connText *widget.Entry
	// Connect button widget
	connBtn *widget.Button
	// Progress bar widget
	progressbar *widget.ProgressBarInfinite
	// Status text label widget
	statusTextLabel *widget.Label
	// Main application tabs widget
	tabs *container.AppTabs
	// libp2p Node
	node *host.Host
	// log text string
	logText string
	// current node Address
	address string
	// Existing peer connections
	connections []*peer.AddrInfo
}

// Start label text
const START_LABEL_TEXT string = "Start new AIP2P node with \"Start\" button."
const STATUS_TEXT_READY = "Ready."
const STATUS_TEXT_STOPPED = "Stopped."
const STATUS_TEXT_STOPPING = "Stopping..."
const STATUS_TEXT_CONNECTING = "Connecting..."
const STATUS_TEXT_CONNECTED = "Connected."
const STATUS_TEXT_DISCONNECTED = "Disconnected."
const STATUS_TEXT_STARTING = "Starting..."
const STATUS_TEXT_STARTED = "Started."
const STATUS_TEXT_ERROR = "Error."

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

		gui.text = widget.NewMultiLineEntry()
		gui.logText = "Welcome to AIP2P Application!\n"
		gui.text.SetText(gui.logText)
		gui.text.Disable()
		gui.text.SetMinRowsVisible(20)
		gui.startBtn = widget.NewButton("Start", func() {
			gui.progressbar.Show()
			gui.progressbar.Start()
			if gui.node != nil {
				gui.statusTextLabel.SetText(STATUS_TEXT_STOPPING)
				err = node.StopListen(gui.node)
				if err != nil {
					gui.LogError(err)
					return
				}
				gui.node = nil
				gui.startBtn.SetText("Start")
				gui.tabs.DisableIndex(1)
				gui.connText.SetText("")
				gui.connText.Disable()
				gui.text.SetText(gui.logText)
				gui.copyBtn.Disable()
				gui.statusTextLabel.SetText(STATUS_TEXT_STOPPED)
				gui.LogInfo(fmt.Sprintf("%s %s", STATUS_TEXT_STOPPED, START_LABEL_TEXT))
				gui.progressbar.Stop()
				gui.progressbar.Hide()
				return
			}
			gui.statusTextLabel.SetText(STATUS_TEXT_STARTING)
			peer, err := node.Listen()
			if err != nil {
				gui.LogError(err)
				return
			}
			addrs, err := node.GetAddress(peer)
			if err != nil {
				gui.LogError(err)
				return
			}
			gui.node = peer
			gui.address = addrs[len(addrs)-1].String()
			gui.startBtn.SetText("Stop")
			gui.tabs.EnableIndex(1)
			gui.connText.Enable()
			gui.LogInfo(fmt.Sprintf("%s Your P2P node address: %v", STATUS_TEXT_STARTED, (*gui.node).Addrs()[len(addrs)-1].String()))
			gui.text.SetText(gui.logText)
			gui.copyBtn.Enable()
			gui.statusTextLabel.SetText(STATUS_TEXT_STARTED)
			gui.progressbar.Stop()
			gui.progressbar.Hide()
		})
		gui.copyBtn = widget.NewButton("Copy Connection String", func() {
			if gui.address != "" {
				gui.window.Clipboard().SetContent(gui.address)
			}
		})
		gui.copyBtn.Disable()
		gui.connectionsTable = widget.NewTable(
			func() (int, int) {
				if len(gui.connections) < 14 {
					return 14, 3
				}
				return len(gui.connections), 3
			},
			func() fyne.CanvasObject { return widget.NewLabel("") }, // Header
			func(tci widget.TableCellID, co fyne.CanvasObject) { // Cell
				if tci.Row == 0 {
					if tci.Col == 0 {
						co.(*widget.Label).SetText("Address")
					} else if tci.Col == 1 {
						co.(*widget.Label).SetText("ID")
					} else if tci.Col == 2 {
						co.(*widget.Label).SetText("Status")
					}
				} else {
					if len(gui.connections) < tci.Row {
						co.(*widget.Label).SetText("")
						return
					}
					if tci.Col == 0 {
						co.(*widget.Label).SetText(gui.connections[tci.Row-1].Addrs[len(gui.connections[tci.Row-1].Addrs)-1].String())
					} else if tci.Col == 1 {
						co.(*widget.Label).SetText(gui.connections[tci.Row-1].ID.String())
					} else if tci.Col == 2 {
						co.(*widget.Label).SetText("Connected")
					}
				}
			})
		gui.connectionsTable.SetColumnWidth(0, 450)
		gui.connectionsTable.SetColumnWidth(1, 450)
		gui.connectionsTable.SetColumnWidth(2, 100)
		gui.connText = widget.NewEntry()
		gui.connText.SetPlaceHolder("Paste Connection String here...")
		gui.connText.Disable()

		gui.connBtn = widget.NewButton("Connect", func() {
			gui.progressbar.Show()
			gui.progressbar.Start()
			gui.connBtn.Disable()
			gui.connText.Disable()
			gui.statusTextLabel.SetText(STATUS_TEXT_CONNECTING)
			gui.LogInfo(fmt.Sprintf("Connecting to %v...", gui.connText.Text))
			peer, err := node.Connect(gui.node, gui.connText.Text)
			if err != nil {
				gui.LogError(err)
				gui.connBtn.Enable()
				gui.connText.Enable()
				return
			}
			gui.connections = append(gui.connections, peer)
			gui.connectionsTable.Refresh()
			gui.statusTextLabel.SetText(STATUS_TEXT_CONNECTED)
			gui.LogInfo(fmt.Sprintf("Connected to %v.", gui.connText.Text))
			gui.connText.SetText("")
			gui.connText.Enable()
			gui.progressbar.Stop()
			gui.progressbar.Hide()
		})
		gui.connBtn.Disable()
		gui.connText.OnChanged = func(text string) {
			if text != "" && gui.node != nil {
				gui.connBtn.Enable()
			} else {
				gui.connBtn.Disable()
			}
		}
		gui.tabs = container.NewAppTabs(
			container.NewTabItem("P2P Node", container.NewVBox(
				gui.text,
				container.NewHBox(
					gui.startBtn,
					gui.copyBtn,
				),
				container.New(
					layout.NewFormLayout(),
					widget.NewLabel("Add new connection:"),
					gui.connText,
				),
				container.NewHBox(
					gui.connBtn,
				),
			)),
			container.NewTabItem("Connections", container.NewMax(gui.connectionsTable)),
		)
		gui.tabs.DisableIndex(1)
		gui.statusTextLabel = widget.NewLabel(STATUS_TEXT_READY)
		gui.progressbar = widget.NewProgressBarInfinite()
		gui.progressbar.Stop()
		gui.progressbar.Hide()
		gui.window.SetContent(container.NewVBox(
			gui.tabs,
			widget.NewSeparator(),
			container.NewHBox(
				gui.statusTextLabel,
				gui.progressbar,
			)),
		)
		gui.window.Resize(fyne.NewSize(1024, 590))
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
		gui.LogInfo(fmt.Sprintf("%s %s", STATUS_TEXT_READY, START_LABEL_TEXT))
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

func (gui *Gui) LogError(err error) {
	gui.statusTextLabel.SetText(STATUS_TEXT_ERROR)
	gui.progressbar.Stop()
	gui.progressbar.Hide()
	gui.logText += fmt.Sprintf("[%s] Error: %v\n", time.Now().Format(time.RFC1123), err)
	gui.text.SetText(gui.logText)
	gui.text.CursorRow = strings.Count(gui.logText, "\n")
}

func (gui *Gui) LogInfo(info string) {
	gui.logText += fmt.Sprintf("[%s] %v\n", time.Now().Format(time.RFC1123), info)
	gui.text.SetText(gui.logText)
	gui.text.CursorRow = strings.Count(gui.logText, "\n")
}
