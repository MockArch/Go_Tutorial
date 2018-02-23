package main 


import (
	"fmt"
	"log"
	"net/http"
	"time"
	"encoding/json"
	"github.com/gorilla/mux"
	)

type Todo struct{
	Name string		
	Compleated bool	
	due  time.time 
}


type  Todos []Todo

func main() {
	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}



func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Todo INDEX!!!")
}


func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{ Name: "Mock" },
		Todo{ Name: "Host Meet Up GO:)" },
	}

	json.NewEncoder(w).Encode(todos)

}

func TodoShow(w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}