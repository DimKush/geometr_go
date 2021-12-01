package service

import (
	"github.com/DimKush/geometry_go/internal/entity/unit"
	"github.com/DimKush/geometry_go/internal/entity/warehouse"
	"github.com/DimKush/geometry_go/internal/repository"
)

type Warehouse interface {
	GetWarehouseById(id int) (*warehouse.ItemWarehouse, error)
}

type Unit interface {
	SetUnit(unit unit.Unit) error
	GetUnit(id int) (*unit.Unit, error)
}

type Service struct {
	Warehouse
	Unit
}

func InitService(repos *repository.Repository) *Service {
	return &Service{
		Warehouse: InitWarehouseService(repos),
		Unit: InitUnitService(repos),
	}
}
