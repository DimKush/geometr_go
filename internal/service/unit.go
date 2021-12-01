package service

import (
	"fmt"

	"github.com/DimKush/geometry_go/internal/entity/unit"
	"github.com/DimKush/geometry_go/internal/repository"
)

type UnitService struct {
	unit_repo repository.Unit
}

func (data *UnitService)SetUnit(unit unit.Unit) error{
	fmt.Println("Unit Service")
	*unit.Geom = string(unit.GeomJson)
	return data.unit_repo.SetUnit(unit)
}

func (data *UnitService)GetUnit(id int) (*unit.Unit, error) {
	return data.unit_repo.GetUnit(id)
}

func InitUnitService(repos *repository.Repository) *UnitService{
	return &UnitService{unit_repo: repos}
}