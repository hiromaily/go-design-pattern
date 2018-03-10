package abstract_factory

import (
	"errors"
	"fmt"
)

//-----------------------------------------------------------------------------
// Car Factory
//-----------------------------------------------------------------------------
const (
	LuxuryCarType   = 1
	FamiliarCarType = 2
)

type CarFactory struct{}

func (c *CarFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamiliarCarType:
		return new(FamiliarCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}

//-----------------------------------------------------------------------------
// Car
//-----------------------------------------------------------------------------
type Car interface {
	GetDoors() int
}

//-----------------------------------------------------------------------------
// Familiar Car
//-----------------------------------------------------------------------------
type FamiliarCar struct{}

func (l *FamiliarCar) GetDoors() int {
	return 5
}
func (l *FamiliarCar) GetWheels() int {
	return 4
}
func (l *FamiliarCar) GetSeats() int {
	return 5
}

//-----------------------------------------------------------------------------
// Luxury Car
//-----------------------------------------------------------------------------
type LuxuryCar struct{}

func (l *LuxuryCar) GetDoors() int {
	return 4
}
func (l *LuxuryCar) GetWheels() int {
	return 4
}
func (l *LuxuryCar) GetSeats() int {
	return 5
}
