package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func Test_Get(t *testing.T){
	//res := httptest.NewRecorder()
	//req,_ := http.NewRequest("GET","/",nil)
	//HomePageHandle(res,req)

	//if res.Code != 200 {
	//	t.Fatalf("Expected status to ==200, but got %d", res.Code)
	//}
	ts := httptest.NewServer(NewRouter())
	defer ts.Close()

	res,_ := http.Get(ts.URL + "/espresso")

	if res.StatusCode !=200 {
		t.Fatalf("Expected status to ==200, but got %d", res.StatusCode)
	}
}

func Test_Post(t *testing.T){
	//res := httptest.NewRecorder()
	//req,_ := http.NewRequest("POST","/",nil)
	//HomePageHandle(res,req)

	//if res.Code == 200 {
	//	t.Fatalf("Expected status to !=200, but got %d", res.Code)
	//}
	
	ts := httptest.NewServer(NewRouter())
	defer ts.Close()

	res,_ := http.Post(ts.URL + "/","",nil)

	if res.StatusCode ==200 {
		t.Fatalf("Expected status to !=200, but got %d", res.StatusCode)
	}
}

