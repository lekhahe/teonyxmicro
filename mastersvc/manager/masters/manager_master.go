package masters

import (
	bu "teonyxmicro/mastersvc/bucontracts"
	bs "teonyxmicro/mastersvc/business"
)

//----------------------------------------
//Create Company
//----------------------------------------
func (m *MasterManager) CreateCompany(company bu.CompanyBO) (uint, error) {

	master := masFac.New(bs.CCompany).(bs.Company)
	masFac.Conn.Begin()
	res, err := master.CreateCompany(company)
	if err != nil {
		masFac.Conn.Rollback()
		return 0, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Update Company
//-----------------------------------------
func (m *MasterManager) UpdateCompany(company bu.CompanyBO) (bool, error) {
	master := masFac.New(bs.CCompany).(bs.Company)
	masFac.Conn.Begin()
	res, err := master.UpdateCompany(company)
	if err != nil {
		masFac.Conn.Rollback()
		return false, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Delete Company
//-----------------------------------------
func (m *MasterManager) DeleteCompany(id uint) (bool, error) {
	master := masFac.New(bs.CCompany).(bs.Company)
	masFac.Conn.Begin()
	res, err := master.DeleteCompany(id)
	if err != nil {
		masFac.Conn.Rollback()
		return false, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Create Address Type
//-----------------------------------------
func (m *MasterManager) CreateAddressType(addressType bu.AddressTypeBO) (uint, error) {
	master := masFac.New(bs.CAddressType).(bs.AddressType)
	masFac.Conn.Begin()
	res, err := master.CreateAddressType(addressType)
	if err != nil {
		masFac.Conn.Rollback()
		return 0, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Update Address Type
//-----------------------------------------
func (m *MasterManager) UpdateAddressType(addressType bu.AddressTypeBO) (bool, error) {
	master := masFac.New(bs.CAddressType).(bs.AddressType)
	masFac.Conn.Begin()
	res, err := master.UpdateAddressType(addressType)
	if err != nil {
		masFac.Conn.Rollback()
		return false, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//----------------------------------------
//Delete Address Type
//----------------------------------------
func (m *MasterManager) DeleteAddressType(id uint) (bool, error) {
	master := masFac.New(bs.CAddressType).(bs.AddressType)
	masFac.Conn.Begin()
	res, err := master.DeleteAddressType(id)
	if err != nil {
		masFac.Conn.Rollback()
		return false, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Get AddressType by Id
//-----------------------------------------
func (m *MasterManager) GetAddressTypeById(id uint) (bu.AddressTypeBO, error) {
	master := masFac.New(bs.CAddressType).(bs.AddressType)
	res, err := master.GetAddressTypeById(id)
	if err != nil {
		return bu.AddressTypeBO{}, err
	}
	return res, nil
}

//-----------------------------------------
//Get Address By Name
//-----------------------------------------
func (m *MasterManager) GetAddressTypeByName(name string) (bu.AddressTypeBO, error) {
	master := masFac.New(bs.CAddressType).(bs.AddressType)
	res, err := master.GetAddressTypeByName(name)
	if err != nil {
		return bu.AddressTypeBO{}, err
	}
	return res, nil
}

//-----------------------------------------
//Get Address
//-----------------------------------------
func (m *MasterManager) GetAllAddressTypes() ([]bu.AddressTypeBO, error) {
	master := masFac.New(bs.CAddressType).(bs.AddressType)
	res, err := master.GetAll()
	if err != nil {
		return []bu.AddressTypeBO{}, err
	}
	return res, nil
}

//-----------------------------------------
//Get Address By Name like
//-----------------------------------------
func (m *MasterManager) GetAllAddressTypeNames(namePart string) ([]bu.AddressTypeBO, error) {
	master := masFac.New(bs.CAddressType).(bs.AddressType)
	res, err := master.GetAllNames(namePart)
	if err != nil {
		return []bu.AddressTypeBO{}, err
	}
	return res, nil
}

//------------------------------------------
//Create Region
//------------------------------------------
func (m *MasterManager) CreateRegion(bo bu.RegionBO) (uint, error) {
	master := masFac.New(bs.CRegion).(bs.Region)
	masFac.Conn.Begin()
	res, err := master.CreateRegion(bo)
	if err != nil {
		masFac.Conn.Rollback()
		return 0, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Update Region
//-----------------------------------------
func (m *MasterManager) UpdateRegion(bo bu.RegionBO) (bool, error) {
	master := masFac.New(bs.CRegion).(bs.Region)
	masFac.Conn.Begin()
	res, err := master.UpdateRegion(bo)
	if err != nil {
		masFac.Conn.Rollback()
		return false, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Delete Region
//-----------------------------------------
func (m *MasterManager) DeleteRegion(id uint) (bool, error) {
	master := masFac.New(bs.CRegion).(bs.Region)
	masFac.Conn.Begin()
	res, err := master.DeleteRegion(id)
	if err != nil {
		masFac.Conn.Rollback()
		return false, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//----------------------------------------
//Get All Region
//----------------------------------------
func (m *MasterManager) GetAllRegion() ([]bu.RegionBO, error) {
	master := masFac.New(bs.CRegion).(bs.Region)
	res, err := master.GetAllRegion()
	if err != nil {
		return []bu.RegionBO{}, err
	}
	return res, nil
}

//-----------------------------------------
//Get Region By Id
//-----------------------------------------
func (m *MasterManager) GetRegionById(id uint) (bu.RegionBO, error) {
	master := masFac.New(bs.CRegion).(bs.Region)
	res, err := master.GetRegionById(id)
	if err != nil {
		return bu.RegionBO{}, err
	}
	return res, nil
}

//-----------------------------------------
//Get Region By Name
//-----------------------------------------
func (m *MasterManager) GetRegionByName(name string) (bu.RegionBO, error) {
	master := masFac.New(bs.CRegion).(bs.Region)
	res, err := master.GetRegionByName(name)
	if err != nil {
		return bu.RegionBO{}, err
	}
	return res, nil
}

//------------------------------------------
//Create State
//------------------------------------------
func (m *MasterManager) CreateState(bo bu.StateBO) (uint, error) {
	master := masFac.New(bs.CState).(bs.State)
	masFac.Conn.Begin()
	res, err := master.CreateState(bo)
	if err != nil {
		masFac.Conn.Rollback()
		return 0, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//------------------------------------------
//Update State
//------------------------------------------
func (m *MasterManager) UpdateState(bo bu.StateBO) (bool, error) {
	master := masFac.New(bs.CState).(bs.State)
	masFac.Conn.Begin()
	res, err := master.UpdateState(bo)
	if err != nil {
		masFac.Conn.Rollback()
		return false, err
	}
	masFac.Conn.Commit()
	return res, nil

}

//-------------------------------------------
//Delete State
//-------------------------------------------
func (m *MasterManager) DeleteState(id uint) (bool, error) {
	master := masFac.New(bs.CState).(bs.State)
	masFac.Conn.Begin()
	res, err := master.DeleteState(id)
	if err != nil {
		masFac.Conn.Rollback()
		return false, err
	}
	masFac.Conn.Commit()
	return res, nil
}

//-------------------------------------------
//Get State by Id
//-------------------------------------------
func (m *MasterManager) GetStateById(id uint) (bu.StateBO, error) {
	master := masFac.New(bs.CState).(bs.State)
	res, err := master.GetStateById(id)
	return res, err
}

//-------------------------------------------
//Get State by Country
//-------------------------------------------
func (m *MasterManager) GetStateByCountryId(id uint) ([]bu.StateBO, error) {
	master := masFac.New(bs.CState).(bs.State)
	res, err := master.GetStateByCountryId(id)
	return res, err
}

//-------------------------------------------
//Get State by name
//-------------------------------------------
func (m *MasterManager) GetStateByName(name string) (bu.StateBO, error) {
	master := masFac.New(bs.CState).(bs.State)
	res, err := master.GetStateByName(name)
	return res, err
}

//-------------------------------------------
//Get All states
//-------------------------------------------
func (m MasterManager) GetAllStates() ([]bu.StateBO, error) {
	master := masFac.New(bs.CState).(bs.State)
	res, err := master.GetAll()
	return res, err
}
