package main

import (
	"encoding/json"
	"fmt"
	"log"

	// "math/rand
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"math/rand"
)
type Movie struct{
 ID string `json:"id"`
 Isbm string `json:"isbm"`
 Title string `json:"title"`
 Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`

}

var movies []Movie
// slice of movie

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(&movies)
}
func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item:=range movies{
		if item.ID==params["id"]{
			movies=append(movies[:index],movies[index+1:]... )	
           break
		}

	}
	json.NewEncoder(w).Encode(movies)

}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params :=mux.Vars(r)
	for _,item :=range movies{
		if(item.ID==params["id"]){
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("user not found")


}
func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_=json.NewDecoder(r.Body).Decode(&movie)
	movie.ID =strconv.Itoa(rand.Intn(10000000))
	movies=append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	for index,item:=range movies{
		if(item.ID==params["id"]){
			movies=append(movies[:index],movies[index+1:]... )
			var newMovie Movie
			_=json.NewDecoder(r.Body).Decode(&newMovie)
			newMovie.ID =params["id"]
			movies=append(movies, newMovie)
			json.NewEncoder(w).Encode(movies)
			break
		}

	}

}
func main(){
	r:=mux.NewRouter()
	movies=append(movies, Movie{
		ID: "1",
		Isbm: "438122",
		Title: "Movie 1",
		Director: &Director{
			Firstname: "john",
			Lastname: "Doe",
		},
	})
	movies=append(movies, Movie{
		ID:"2",
	    Isbm: "34550",
		Title: "black panther",
		Director: &Director{
			Firstname: "chadwick",
			Lastname: "bosman",
		},

	})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")
	fmt.Printf("started server at port 8080")
	if err:=http.ListenAndServe(":8080",r); err!=nil{
		log.Fatal(err)
	}
	
}