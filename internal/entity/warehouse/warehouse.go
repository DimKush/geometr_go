package warehouse

import (
	"encoding/json"
	"time"
)

type ItemWarehouse struct {
	Id          int64           `json:"id" gorm:"id"`
	CreatedAt   time.Time       `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"updated_at"`
	DeletedAt   time.Time       `json:"deleted_at" gorm:"deleted_at"`
	Name        string          `json:"name" gorm:"name"`
	Description string          `json:"description" gorm:"description"`
	Latitude    float64         `json:"latitude" gorm:"latitude"`
	Longitude   float64         `json:"longitude" gorm:"longitude"`
	Geom        json.RawMessage `json:"geometry"`
	Poly        json.RawMessage `json:"polygon" gorm:""`
	MultiPoly   json.RawMessage `json:"multi_polygon" gorm:"multi_poly"`
}
