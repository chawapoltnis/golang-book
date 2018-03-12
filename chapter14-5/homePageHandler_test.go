package home

import (
	"strings"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req,_ := http.NewRequest("POST","/json", strings.NewReader(`{"first_name":"espresso","last_name":"longshot"}`))
	home := HomePageHandler{}
	home.ServeHTTP(res,req)

	if res.Code != 201 {
		t.Fatalf("Expected status to ==201, but got %d",res.Code)
	}

	user := new(User)
	json.NewDecoder(res.Body).Decode(user)

	if user.FirstName != "espresso" {
		t.Fatalf("Expectes user firstname to== espresso, but go %s",user.FirstName)
	}
}