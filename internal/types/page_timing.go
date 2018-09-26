package types

type PageTiming struct {
	OnContentLoad int    `json:"onContentLoad"`
	OnLoad        int    `json:"onLoad"`
	Comment       string `json:"comment,omitempty"`
}
