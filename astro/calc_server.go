package main
import (
	"fmt";
	"net/http";
	"os";
	"encoding/json"
)


func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "You've browsed to the right place, now try accessing /rocket-equation", r.URL.Path[1:])
}


func tsiolkovsky_handler(w http.ResponseWriter, r *http.Request){
		
	names := [10] string {"Dev", "Jeff", "Sion", "Asok", "Gwen", "Turbid Tony", "Ahmed", "Jewell", "Wench Five", "Stryder" }
	points := [] *user.Point { 
		&user.Point{Value: 0, Reason: "What a great guy"},
	}
	for i :=0 ;i<10 ; i++{
		users[i] = user.User{
			Username: names[i],
			Points: points,
		}
	}
	b,_ := json.Marshal(users)
	fmt.Fprintf(w,string(b))
}

func about_handler(w http.ResponseWriter, r *http.Request){
}


func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/rocket-equation", tsiolkovsky_handler)
	port := os.Getenv("PORT")
	fmt.Println("Starting on port",port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
