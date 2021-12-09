package service

import (
	"encoding/json"
	"fmt"

	"github.com/spatial-go/geoos/geojson"
	"github.com/spatial-go/geoos/planar"
	"github.com/spatial-go/geoos/space"

	"github.com/DimKush/geometry_go/internal/entity/unit"
	"github.com/DimKush/geometry_go/internal/repository"
)

type UnitService struct {
	unit_repo repository.Unit
}

func (data *UnitService)SetUnit(unit unit.Unit) error{
	fmt.Println("Unit Service")

	unit.Geom = []byte(unit.JsonZone)

	return data.unit_repo.SetUnit(unit)
}

func (data *UnitService)IntersectUnits(first int, second int) (bool, *[]byte, error){
	unitOne, err := data.unit_repo.GetUnit(first)
	if err != nil {
		fmt.Printf("Error 1 %v\n", err)
		return false,nil, err
	}
	fmt.Printf("\n%d : %v", first, unitOne)

	unitTwo, err := data.unit_repo.GetUnit(second)
	if err != nil {
		fmt.Printf("Error 2 %v\n", err)
		return false,nil, err
	}
	fmt.Printf("\n%d : %v", second, unitTwo)

	fcFirst := geojson.NewFeatureCollection()
	err  = json.Unmarshal(unitOne.Geom, &fcFirst)
	if err != nil {
		fmt.Printf("Error 3 %v\n", err)
		return false,nil, err
	}


	fcSecond := geojson.NewFeatureCollection()
	err  = json.Unmarshal(unitTwo.Geom, &fcSecond)
	if err != nil {
		fmt.Printf("Error 3 %v\n", err)
		return false,nil, err
	}
	fmt.Printf("\n%p", fcSecond)

	polyOne := fcFirst.Features[0].Geometry.Coordinates.(space.Polygon)
	polyTwo := fcSecond.Features[0].Geometry.Coordinates.(space.Polygon)

	strg := planar.NormalStrategy()
	got, err := strg.Intersection(polyOne, polyTwo)
	if err != nil {
		fmt.Printf("Error 3 %v\n", err)
		return false,nil, err
	}

	bytes, err := fcFirst.Features[0].Geometry.MarshalJSON()
	if got == nil {
		return false, &bytes, nil
	} else {
		return true, &bytes, nil
	}
}

func (data *UnitService)GetUnit(id int) (*unit.Unit, error) {
	unitOut, err := data.unit_repo.GetUnit(id)
	if err != nil {
		return nil, err
	}

	unitOut.JsonZone = json.RawMessage(unitOut.Geom)

	return unitOut, nil
}

func InitUnitService(repos *repository.Repository) *UnitService{
	return &UnitService{unit_repo: repos}
}