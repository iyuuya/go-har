package types

type Page struct {
	StartedDateTime string     `json:"startedDateTime"`
	ID              string     `json:"id"`
	Title           string     `json:"title"`
	PageTiming      PageTiming `json:"pageTiming"`
	Comment         string     `json:"comment,omitempty"`
}
