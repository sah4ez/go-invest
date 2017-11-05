package moex

type Cursor struct {
	Metadata MetadataCursor `json:"metadata"`
	Columns  []string       `json:"columns"`
	Data     [][]int64      `json:"data"` // [[start, size_all, limit]]
}

type MetadataCursor struct {
	INDEX    MetaValue `json:"INDEX"`
	TOTAL    MetaValue `json:"TOTAL"`
	PAGESIZE MetaValue `json:"PAGESIZE"`
}

// Next retrun next position cursor and Has next iteration or no
func (c *Cursor) Next() (int64, bool) {
	next := c.Data[0][0] + c.Data[0][2]
	if next < c.Data[0][1] {
		return next, true
	}
	return c.Data[0][0], false

}
