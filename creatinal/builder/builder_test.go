package builder_test

import (
	. "github.com/hiromaily/go-design-pattern/creatinal/builder"
	tu "github.com/hiromaily/golibs/testutil"
	"testing"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[Builder]")
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func testCar(t *testing.T, director ManufacturingDirector) {
	carBuilder := &CarBuilder{}
	director.SetBuilder(carBuilder)
	director.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}
}

func testBike(t *testing.T, director ManufacturingDirector) {
	bikeBuilder := &BikeBuilder{}

	director.SetBuilder(bikeBuilder)
	director.Construct()

	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = 1

	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n", motorbike.Wheels)
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s\n", motorbike.Structure)
	}
}

func testBus(t *testing.T, director ManufacturingDirector) {
	busBuilder := &BusBuilder{}

	director.SetBuilder(busBuilder)
	director.Construct()

	bus := busBuilder.GetVehicle()

	if bus.Wheels != 8 {
		t.Errorf("Wheels on a bus must be 8 and they were %d\n", bus.Wheels)
	}

	if bus.Structure != "Bus" {
		t.Errorf("Structure on a bus must be 'Bus' and was %s\n", bus.Structure)
	}

	if bus.Seats != 30 {
		t.Errorf("Seats on a bus must be 30 and they were %d\n", bus.Seats)
	}

}

func TestBuilder(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	// Car
	testCar(t, manufacturingComplex)

	// Bike
	testBike(t, manufacturingComplex)

	// Bus
	testBus(t, manufacturingComplex)
}
