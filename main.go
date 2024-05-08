package main
import (
	"fmt"
	"log"
	"encoding/json"
	// "math/rand
	"net/http"
	// "strconv"
	"github.com/gorilla/mux"

)
type Movie struct{
 ID string `json:"id"`
 Isbm string `json:"isbm"`
 title string `json:"title"`
 Director *Director `json:"director"`
}

type Director struct{
	firstname string `json:"firstname"`
	lastname string `json:"lastname"`

}
var movies []Movie
// slice of movie

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item:=range movies{
		if item.ID==params["id"]{

		}

	}
	json.NewEncoder().Encode()

}
func main(){
	r:=mux.NewRouter()
	movies=append(movies, Movie{
		ID: "1",
		Isbm: "438122",
		title: "Movie 1",
		Director: &Director{
			firstname: "john",
			lastname: "Doe",
		},
	})
	movies=append(movies, Movie{
		ID:"2",
	    Isbm: "34550",
		title: "black panther",
		Director: &Director{
			firstname: "chadwick",
			lastname: "bosman",
		},

	})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	// r.HandleFunc("/movies/:id",getMovie).Methods("GET")
	// r.HandleFunc("/movies",createMovie).Methods("POST")
	// r.HandleFunc("/movies/{id}",updateMovie).Methods("UPDATE")
	// r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")
	fmt.Printf("started server at port 8080")
	if err:=http.ListenAndServe(":8080",r); err!=nil{
		log.Fatal(err)
	}
	
}