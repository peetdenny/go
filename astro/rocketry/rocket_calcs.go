package rocketry 
import (
	"math";
)

const grabbity = 9.81

// dv = Gravity * impulse * natural log of (starting mass / end mass)
func Calc_delta_v(mass1 float64, mass2 float64, impulse float64 ) float64{
	return grabbity * impulse * math.Log(mass1/mass2)
}
