package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func main() {
	//Uploadobject("hello", "jenkins19","now")
	//Listobject("jenkins19")
	//Downloadobj("now","jenkins19")
	//Deleteobj("now","jenkins19")
	router := mux.NewRouter()
	router.HandleFunc("/uploadobj" ,Uploadobject).Methods("PUT")//(/uploadobj?filename=&bucketname=&objectname=)
	router.HandleFunc("/listobj" ,Listobject).Methods("GET") //(/listobj?bucketname=)
	router.HandleFunc("/downloadobj" ,Downloadobj).Methods("GET") //(/downloadobj?bucketname=&objectname=)
	router.HandleFunc("/deleteobj" ,Deleteobj).Methods("DELETE") //(/deleteobj?bucketname=&objectname=)
	log.Fatal(http.ListenAndServe(":8081", router))
}

