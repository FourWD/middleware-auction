package orm

import "github.com/FourWD/middleware-auction/model"

type VehicleFinance0003 struct {
	model.VehicleModel
}

func (VehicleFinance0003) TableName() string {
	return "vehicle-finances-0003"
}
