package types

type Cache struct {
	BeforeRequest CacheEntry `json:"beforeRequest,omitempty"`
	AfterRequest  CacheEntry `json:"afterRequest,omitempty"`
	Comment       string     `json:"comment,omitempty"`
}

type CacheEntry struct {
	Expires    string `json:"expires,omitempty"`
	LastAccess string `json:"lastAccess"`
	ETag       string `json:"eTag"`
	HitCount   int    `json:"hitCount"`
	Comment    string `json:"comment,omitempty"`
}
