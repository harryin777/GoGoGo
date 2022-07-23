package construct

// Car 汽车类
type Car struct {
	Brand    string `json:"brand"`
	Turbo    bool   `json:"turbo"`
	MaxSpeed int    `json:"max_speed"`
}

type CarOpt func(*Car)

func SetBrand(brand string) CarOpt {
	return func(car *Car) {
		if len(brand) != 0 {
			car.Brand = brand
		}
	}
}

func SetTurbo(turbo bool) CarOpt {
	return func(car *Car) {
		if turbo {
			car.Turbo = turbo
		}
	}
}

func SetMaxSpeed(maxSpeed int) CarOpt {
	return func(car *Car) {
		if maxSpeed != 0 {
			car.MaxSpeed = maxSpeed
		}
	}
}

func NewCar(opts ...CarOpt) *Car {
	car := &Car{
		Brand:    "",
		Turbo:    false,
		MaxSpeed: 0,
	}
	for _, opt := range opts {
		opt(car)
	}
	return car
}
