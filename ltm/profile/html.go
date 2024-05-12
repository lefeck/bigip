package profile

import "github.com/lefeck/go-bigip"

type HtmlConfigList struct {
	Items    []HtmlConfig `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

type HtmlConfig struct {
	Kind             string   `json:"kind"`
	Name             string   `json:"name"`
	Partition        string   `json:"partition"`
	FullPath         string   `json:"fullPath"`
	Generation       int      `json:"generation"`
	SelfLink         string   `json:"selfLink"`
	AppService       string   `json:"appService"`
	ContentDetection string   `json:"contentDetection"`
	ContentSelection []string `json:"contentSelection"`
	DefaultsFrom     string   `json:"defaultsFrom"`
	Description      string   `json:"description"`
}

const HtmlEndpoint = "html"

type HtmlResource struct {
	b *bigip.BigIP
}