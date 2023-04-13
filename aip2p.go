package aip2p

import (
	"webimizer.dev/aip2p/app"
)

func StartApplicationGUI() {
	gui := new(app.Gui)
	gui.Start()
}
