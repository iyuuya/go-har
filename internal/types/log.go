package types

type Log struct {
	Version string  `json:"version"`
	Creator Creator `json:"creator"`
	Browser Browser `json:"browser,omitempty"`
	Pages   []Page  `json:"pages,omitempty"`
	Entries []Entry `json:"entries"`
	Comment string  `json:"comment,omitempty"`
}
