package home

import (
	"encoding/json"
	"time"
	"net/http"
)

type HomePageHandler struct{}
func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	user := new(User)
	json.NewDecoder(r.Body).Decode(user)
	user.CreatedAt = time.Now()

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)

	data,_ := json.Marshal(user)
	w.Write(data)
}

func main() {
	http.Handle("/",&HomePageHandler{})

	http.ListenAndServe(":3000",nil)
}

type User struct{
	FirstName string	`json:"first_name"`
	LastName string
	Email string
	CreatedAt time.Time
}