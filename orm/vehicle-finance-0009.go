package orm

import "github.com/FourWD/middleware-auction/model"

type VehicleFinance0009 struct {
	model.VehicleModel
}

func (VehicleFinance0009) TableName() string {
	return "vehicle-finances-0009"
}
