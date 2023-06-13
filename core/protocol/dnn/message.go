package dnn

type Message struct {

	// Type of message
	Type string `json:"type"`
	Data string `json:"data"`
}

const TYPE_NODELIST = "nodelist"
