package types

type PostData struct {
	MimeType string  `json:"mimeType"`
	Params   []Param `json:"params"`
	Text     string  `json:"text"`
	Comment  string  `json:"comment,omitempty"`
}

type Param struct {
	Name        string `json:"name"`
	Value       string `json:"value,omitempty"`
	FileName    string `json:"fileName,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Comment     string `json:"comment,omitempty"`
}
