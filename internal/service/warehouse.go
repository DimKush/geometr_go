package service

import (
	"github.com/DimKush/geometry_go/internal/entity/warehouse"
	"github.com/DimKush/geometry_go/internal/repository"
)

type WarehouseService struct {
	whs_repo repository.Warehouse
}

func (data *WarehouseService)GetWarehouseById(id int) (*warehouse.ItemWarehouse, error){
	whs, err := data.whs_repo.GetWarehouseById(id)
	if err != nil {
		return nil, err
	}
	return whs, nil
}

func InitWarehouseService(repos *repository.Repository) *WarehouseService{
	return &WarehouseService{whs_repo: repos}
}