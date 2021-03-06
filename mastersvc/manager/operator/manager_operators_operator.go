package operator

import (
	bu "teonyxmicro/mastersvc/bucontracts"
	bs "teonyxmicro/mastersvc/business"
)

//----------------------------------------------
//Create operator
//----------------------------------------------
func (o *OprManager) CreateOperator(bo bu.OperatorBO) (uint, error) {

	op := OpFac.New(bs.CVhOperator).(bs.Operator)
	OpFac.Conn.Begin()
	res, err := op.CreateOperator(bo)
	if err != nil {
		OpFac.Conn.Rollback()
		return res, err
	}
	OpFac.Conn.Commit()
	return res, nil
}

//----------------------------------------------
//Update operator
//----------------------------------------------
func (o *OprManager) UpdateOperator(bo bu.OperatorBO) (bool, error) {
	op := OpFac.New(bs.CVhOperator).(bs.Operator)
	OpFac.Conn.Begin()
	res, err := op.UpdateOperator(bo)
	if err != nil {
		OpFac.Conn.Rollback()
		return res, err
	}
	OpFac.Conn.Commit()
	return res, nil
}

//----------------------------------------------
//Delete operator
//----------------------------------------------
func (o *OprManager) DeleteOperator(id uint) (bool, error) {
	op := OpFac.New(bs.CVhOperator).(bs.Operator)
	OpFac.Conn.Begin()
	res, err := op.DeleteOperator(id)
	if err != nil {
		OpFac.Conn.Rollback()
		return res, err
	}
	OpFac.Conn.Commit()
	return res, nil
}

//---------------------------------------------
//Get operator by Id
//---------------------------------------------
func (o *OprManager) GetOperatorById(id uint) (bu.OperatorBO, error) {
	op := OpFac.New(bs.CVhOperator).(bs.Operator)
	res, err := op.GetOperatorById(id)
	return res, err
}

//---------------------------------------------
//Get operator by vehicleid
//---------------------------------------------
func (o *OprManager) GetOperatorsByVehicleId(id uint) ([]bu.OperatorBO, error) {
	op := OpFac.New(bs.CVhOperator).(bs.Operator)
	res, err := op.GetOperatorsByVehicleId(id)
	return res, err
}
