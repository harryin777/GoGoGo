package construct

import (
	"fmt"
	"testing"
)

func Test_GetNewCar(t *testing.T) {
	car := NewCar(SetBrand("lalala"), SetTurbo(true), SetMaxSpeed(999))
	fmt.Println(*car)
}
