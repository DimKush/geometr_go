package repository

import (
	"fmt"

	"github.com/twpayne/go-geom/encoding/ewkb"
	"github.com/twpayne/go-geom/encoding/geojson"
	"gorm.io/gorm"

	warehouse "github.com/DimKush/geometry_go/internal/entity/warehouse"
)

type WarehouseRep struct {
	db *gorm.DB
}

func (data *WarehouseRep)GetWarehouseById(id int) (*warehouse.ItemWarehouse, error){
	//rows, err := data.db.Table(warehouses).Select("id, created_at, updated_at, deleted_at, name, description, latitude, longitude, st_asewkb(geom)").
	//	Where("id = ?", id).Debug().Rows()
	rows, err := data.db.Table(warehouses).Select("id, name, st_asewkb(geom), st_asewkb(poly), st_asewkb(multi_poly)").
		Where("id = ?", id).Debug().Rows()
	if err != nil {
		fmt.Println("Error : %v", err)
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
		var ewkbPoint ewkb.Point
		var ewkbPoly ewkb.Polygon
		var ewkbMultiPoly ewkb.MultiPolygon

		if err := rows.Scan(&id, &name, &ewkbPoint, &ewkbPoly, &ewkbMultiPoly); err != nil {
			return nil, err
		}

		geometry, err := geojson.Marshal(ewkbPoint.Point)
		if err != nil {
			return nil, err
		}

		polygon, err := geojson.Marshal(ewkbPoly.Polygon)
		if err != nil {
			return nil, err
		}

		multiPolygon, err := geojson.Marshal(ewkbMultiPoly.MultiPolygon)
		if err != nil {
			return nil, err
		}

		return &warehouse.ItemWarehouse{Id: int64(id), Name: name, Geom: geometry, Poly: polygon, MultiPoly: multiPolygon}, nil
	}

	return nil, nil
}

func InitWarehouseRep(database *gorm.DB) *WarehouseRep{
	return &WarehouseRep{db : database}
}


