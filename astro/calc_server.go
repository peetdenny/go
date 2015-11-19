package main
import (
	"fmt";
	"net/http";
	"os";
	"strconv";
	"encoding/json";
	"github.com/peetdenny/go/astro/rocketry";
)

type DeltaVCalc struct {
	Starting_Mass float64
	End_Mass float64
	Impulse float64
	DeltaV float64
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Best Boy Electric Delta-V calculator! Now try accessing /delta-v-calc?starting-mass=2400&end-mass=400&impulse=210. Then try tweaking the request parameters" )
}


func tsiolkovsky_handler(w http.ResponseWriter, r *http.Request){
	// input the starting mass (in kg), the final mass, plus the motor impulse (in seconds), and this returns the delta-v (in m/s)
	values := r.URL.Query()
	starting_mass,_ := strconv.ParseFloat(values.Get("starting-mass"),64)
	end_mass,_ := strconv.ParseFloat(values.Get("end-mass"),64)
	impulse,_ := strconv.ParseFloat(values.Get("impulse"),64)
	dv :=rocketry.Calc_delta_v(starting_mass, end_mass,impulse)
	fmt.Println("Calculated DV as", dv, "from start mass of", starting_mass, "and end mass of", end_mass)
	resp := &DeltaVCalc{
		Starting_Mass: starting_mass, 
		End_Mass: end_mass,
		Impulse: impulse,
		DeltaV: dv,
	}
	b,_ := json.Marshal(resp)
	fmt.Fprintf(w,string(b))
}

func about_handler(w http.ResponseWriter, r *http.Request){
}


func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/delta-v-calc", tsiolkovsky_handler)
	port := os.Getenv("PORT")
	fmt.Println("Starting on port",port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
