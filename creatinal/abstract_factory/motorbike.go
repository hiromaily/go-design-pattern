package abstract_factory

import (
	"errors"
	"fmt"
)

//-----------------------------------------------------------------------------
// Motorbike Factory
//-----------------------------------------------------------------------------
const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}

//-----------------------------------------------------------------------------
// Motorbike
//-----------------------------------------------------------------------------
type Motorbike interface {
	GetType() int
}

//-----------------------------------------------------------------------------
// Sport Motorbike
//-----------------------------------------------------------------------------
type SportMotorbike struct{}

func (s *SportMotorbike) GetWheels() int {
	return 2
}
func (s *SportMotorbike) GetSeats() int {
	return 1
}
func (s *SportMotorbike) GetType() int {
	return SportMotorbikeType
}

//-----------------------------------------------------------------------------
// Cruise Motorbike
//-----------------------------------------------------------------------------
type CruiseMotorbike struct{}

func (c *CruiseMotorbike) GetWheels() int {
	return 2
}
func (c *CruiseMotorbike) GetSeats() int {
	return 2
}
func (c *CruiseMotorbike) GetType() int {
	return CruiseMotorbikeType
}
