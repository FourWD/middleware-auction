package orm

import "middleware-auction/model"

type VehicleFinance0002 struct {
	model.VehicleModel
}

func (VehicleFinance0002) TableName() string {
	return "vehicle-finances-0002"
}
