package orm

import "middleware-auction/model"

type VehicleFinance0004 struct {
	model.VehicleModel
}

func (VehicleFinance0004) TableName() string {
	return "vehicle-finances-0004"
}
