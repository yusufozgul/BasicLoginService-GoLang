package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "./Models"
)

func main() {
	RunServer()
}

func RunServer() {
	http.HandleFunc("/login", HandleLogin)
	http.ListenAndServe(":8282", nil)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		HandleDocumantationPage(w, r)
		return
	} else if r.Method != "POST" {

		return
	}

	w.Header().Add("Content-Type", "application/json")

	var responseData LoginResponse
	var login LoginPostData

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(body, &login)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if login.Username == "" || login.Password == "" {
		w.WriteHeader(400)
		return
	}

	if (login.Username == "mobil") && (login.Password == "mobiluygulamagelistirme") {

		personData := PersonData{Name: "Mobil", Surname: "Programlama", City: "Manisa"}
		locationData := Location{Lat: 38.5002, Lon: 27.7084}

		responseData = LoginResponse{
			Message: "Başarılı",
			Success: true,
			Data: &LoginData{
				Person:   personData,
				Location: locationData}}
	} else {
		responseData = LoginResponse{Message: "Kullanıcı adı veya şifre hatalı", Success: false}
		w.WriteHeader(401)
	}

	data, _ := json.Marshal(responseData)
	w.Write(data)
}

func HandleDocumantationPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://yusufozgul.com/devs/mobilProje/doc/", 301)
}

type LoginPostData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
