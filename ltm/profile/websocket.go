package profile

import "github.com/lefeck/go-bigip"

type WebSocketList struct {
	Items    []WebSocket `json:"items,omitempty"`
	Kind     string      `json:"kind,omitempty"`
	SelfLink string      `json:"selflink,omitempty"`
}

type WebSocket struct {
	Kind                   string `json:"kind"`
	Name                   string `json:"name"`
	Partition              string `json:"partition"`
	FullPath               string `json:"fullPath"`
	Generation             int    `json:"generation"`
	SelfLink               string `json:"selfLink"`
	AppService             string `json:"appService"`
	CompressMode           string `json:"compressMode"`
	Compression            string `json:"compression"`
	DefaultsFrom           string `json:"defaultsFrom"`
	Description            string `json:"description"`
	Masking                string `json:"masking"`
	NoDelay                string `json:"noDelay"`
	PayloadProcessingMode  string `json:"payloadProcessingMode"`
	PayloadProtocolProfile string `json:"payloadProtocolProfile"`
	WindowBits             int    `json:"windowBits"`
}

const WebSocketEndpoint = "websocket"

type WebSocketResource struct {
	b *bigip.BigIP
}
