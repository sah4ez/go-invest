package moex

type MetaValue struct {
	Type    string `json:"type,omitempty"`
	Bytes   int64  `json:"bytes,omitempty"`
	MaxSize int64  `json:"max_size,omitempty"`
}
