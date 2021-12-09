package unit

import "encoding/json"

type Unit struct {
	Id       int             `json:"id"`
	Name     string          `json:"name"`
	JsonZone json.RawMessage `json:"zone" gorm:"-"`
	Geom     []byte          `json:"-" gorm:"geom"`
}
