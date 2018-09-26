package types

type Entry struct {
	Pageref         string   `json:"pageref,omitempty"`
	StartedDateTime string   `json:"startedDateTime"`
	Time            float32  `json:"time"`
	Request         Request  `json:"request"`
	Response        Response `json:"response"`
	Cache           Cache    `json:"cache"`
	Timings         Timings  `json:"timings"`
	ServerIPAddress string   `json:"serverIPAddress,omitempty"`
	Connection      string   `json:"connection,omitempty"`
	Comment         string   `json:"comment,omitempty"`
}
