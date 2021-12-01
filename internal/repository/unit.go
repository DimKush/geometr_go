package repository

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"

	"github.com/DimKush/geometry_go/internal/entity/unit"
)

type UnitRep struct {
	db *gorm.DB
}


func (data *UnitRep)SetUnit(unit unit.Unit) error {
	fmt.Printf("\nUnit: %v\n",unit)

	err := data.db.Table(units).Debug().Create(&unit).Error
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (data *UnitRep)GetUnit(id int) (*unit.Unit, error) {
	rows, err := data.db.Table(units).Select("id, name, geom").
		Where("id = ?", id).Rows()
	if err != nil {
		fmt.Println("Error : %v", err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	for rows.Next(){
		var id int
		var name string
		var geomPointStr json.RawMessage

		if err := rows.Scan(&id, &name, &geomPointStr); err != nil {
			return nil, err
		}

		if err != nil {
			return nil, err
		}

		return &unit.Unit{Id: id, Name: name, GeomJson: geomPointStr}, nil
	}

	return nil, nil
}


func InitUnitRep(database *gorm.DB) *UnitRep{
	return &UnitRep{db : database}
}