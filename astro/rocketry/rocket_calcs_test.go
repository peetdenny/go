package rocketry 

import( 
	"fmt";
	"testing";
)

func fuzzy_match(l float64, r float64) bool{
	//The canonical use of math/big.NewFloat doesn't seem to be available in this version. Hence rolling my own dubious comparator, to compare to one decimal place
	L := int32((l*10))
	R := int32((r*10))
	return L == R
}

func TestIdealRocketEquation(t *testing.T){
	fmt.Println("Testing")
	mass1 := 2400.0
	mass2 := 400.0
	specificImpulse := 210.0
	dv := Calc_delta_v(mass1, mass2, specificImpulse)
	fmt.Println("DV calculated as", dv)	
	if ! fuzzy_match(dv, 3691.2) {
		t.Error("For", mass1, mass2, "and", specificImpulse,
			"Expected", 3691.2,
			"Instead, received", dv,

		)
	}
}
