package types

type Request struct {
	Method      string   `json:"method"`
	URL         string   `json:"url"`
	HTTPVersion string   `json:"httpVersion"`
	Cookies     []Cookie `json:"cookies"`
	Headers     []NVP    `json:"headers"`
	QueryString []NVP    `json:"queryString"`
	PostData    PostData `json:"postData,omitempty"`
	HeaderSize  int      `json:"headerSize"`
	BodySize    int      `json:"bodySize"`
	Comment     string   `json:"comment,omitempty"`
}
