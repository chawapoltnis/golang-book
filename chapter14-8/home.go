package main

import (
	"time"
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
	r.HandleFunc("/users",UsersHandle).Methods("GET")
	r.HandleFunc("/{name}",HomePageHandle).Methods("GET")
	r.Use(loggingMiddleware)	//do everytime in handlefunc
	return r
}

func loggingMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		start:=time.Now()
		fmt.Printf("Start at %v\n",start)

		next.ServeHTTP(w,r)		//do everytime when run handlefunc

		fmt.Printf("Completed in %v\n", time.Since(start))
	})
}

