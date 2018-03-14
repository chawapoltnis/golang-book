package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func HomePageHandle(w http.ResponseWriter, r *http.Request) {
	//name := r.URL.Query().Get("name")
	//if name == "" {
	//	name = "World"
	//}
	//fmt.Fprintf(w, "Hello, %s!",name)
	vars := mux.Vars(r)
	name := vars["name"]

	if vars["name"]==""{
		name ="World"
	}
	fmt.Fprintf(w, "Hello, %s!",name)
}

func UsersHandle(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Users Page")
}

func main() {
	//http.HandleFunc("/",HomePageHandle)
	//http.HandleFunc("/users",UsersHandle)
	//http.ListenAndServe(":3000",nil)

	http.ListenAndServe(":3000",NewRouter())
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/{name}",HomePageHandle).Methods("GET")
	r.HandleFunc("/users",UsersHandle).Methods("GET")
	return r
}

