package bins

import (
	"encoding/json"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

type Bins struct {
	Bins []Bin `json:"bins"`
}

func (bins Bins) ToBytes() []byte {
	bytes, err := json.Marshal(bins)
	if err != nil {
		panic(err)
	}
	return bytes
}
