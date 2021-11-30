package service

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/DimKush/geometry_go/internal/entity/unit"
	"github.com/DimKush/geometry_go/internal/repository"
)

type UnitService struct {
	unit_repo repository.Unit
}

// JSONB Interface for JSONB Field of yourTableName Table
type JSONB []interface{}

// Value Marshal
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(b,&a)
}

func (data *UnitService)SetUnit(unit unit.Unit) error{
	fmt.Println("Unit Service")
	return data.unit_repo.SetUnit(unit)
}

func InitUnitService(repos *repository.Repository) *UnitService{
	return &UnitService{unit_repo: repos}
}