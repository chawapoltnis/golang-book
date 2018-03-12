package home

import (
	"net/http"
	"fmt"
)

func HomePageHandler(res http.ResponseWriter, req *http.Request){
	//res.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(res, "Hello, World!")
}