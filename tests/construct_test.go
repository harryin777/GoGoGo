package tests

import (
	"fmt"
	"test1/construct"
	"testing"
)

func Test_GetNewCar(t *testing.T) {
	car := construct.NewCar(construct.SetBrand("lalala"), construct.SetTurbo(true), construct.SetMaxSpeed(999))
	fmt.Println(*car)
}
