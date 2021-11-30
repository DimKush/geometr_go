package repository

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/DimKush/geometry_go/internal/entity/unit"
)

type UnitRep struct {
	db *gorm.DB
}


func (data *UnitRep)SetUnit(unit unit.Unit) error {
	fmt.Println("Unit")
	err := data.db.Table(units).Create(&unit).Error
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}


func InitUnitRep(database *gorm.DB) *UnitRep{
	return &UnitRep{db : database}
}