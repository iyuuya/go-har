package types

type Response struct {
	Status      int      `json:"status"`
	StatusText  string   `json:"statusText"`
	HTTPVersion string   `json:"httpStatus"`
	Cookies     []Cookie `json:"cookies"`
	Headers     []NVP    `json:"headers"`
	Content     Content  `json:"content"`
	RedirectURL string   `json:"redirectURL"`
	HeadersSize int      `json:"headersSize"`
	BodySize    int      `json:"bodySize"`
	Comment     string   `json:"comment,omitempty"`
}
