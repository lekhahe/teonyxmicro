package vehicles

import (
	"log"
	bu "teonyxmicro/mastersvc/bucontracts"
	bs "teonyxmicro/mastersvc/business"
	ma "teonyxmicro/mastersvc/manager"
)

var vehicleFac *bs.RuleFactory

func init() {
	conn, err := ma.Setconn()
	if err != nil {
		log.Fatal("error in master manager initialisation")
	}
	vehicleFac = &bs.RuleFactory{Conn: conn}
}

type IVehicleManager interface {
	CreateVehicle(vehicle bu.VehicleBO) (uint, error)
	UpdateVehicle(vehicle bu.VehicleBO) (bool, error)
	DeleteVehicle(vehicleId uint) (bool, error)
	GetVehicleById(vehicleId uint) (bu.VehicleBO, error)
	GetVehicleByRegistration(registration string) (bu.VehicleBO, error)
	GetVehiclesByFleetId(fleetId uint) ([]bu.VehicleBO, error)
	CreateVehicleHistory(history bu.VehicleHistoryBO) (uint, error)
	UpdateVehicleHistory(history bu.VehicleHistoryBO) (bool, error)
	DeleteVehicleHistory(id uint) (bool, error)
	GetVehicleHistoryByVehicleId(vehicleId uint) ([]bu.VehicleHistoryBO, error)
	CreateVehicleLocation(ad bu.VehicleAddressBO) (uint, error)
	UpdateVehicleLocation(ad bu.VehicleAddressBO) (bool, error)
	DeleteVehicleLocation(id uint) (bool, error)
	GetVehicleLocationByVehicle(vehicleId uint) ([]bu.VehicleAddressBO, error)
	CreateVehicleMake(bo bu.VehicleMakeBO) (uint, error)
	UpdateVehicleMake(bo bu.VehicleMakeBO) (bool, error)
	DeleteVehicleMake(id uint) (bool, error)
	GetAllVehicleMake() ([]bu.VehicleMakeBO, error)
	GetVehicleMakeById(id uint) (bu.VehicleMakeBO, error)
}

type VehicleManager struct{}

func New() *VehicleManager {
	return &VehicleManager{}
}
