package types

type Timings struct {
	Blocked float32 `json:"blocked,omitempty"`
	DNS     float32 `json:"dns,omitempty"`
	Connect float32 `json:"connect,omitempty"`
	Send    float32 `json:"send"`
	Wait    float32 `json:"wait"`
	Receive float32 `json:"receive"`
	SSL     float32 `json:"ssl,omitempty"`
	Comment string  `json:"comment,omitempty"`
}
