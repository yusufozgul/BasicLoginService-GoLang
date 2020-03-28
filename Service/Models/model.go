package Models

type LoginResponse struct {
	Error   error     `json:"error"`
	Message string    `json:"message"`
	Success bool      `json:"success"`
	Data    LoginData `json:"data"`
}

type LoginData struct {
	Person   PersonData `json:"profile"`
	Location Location   `json:"location"`
}

type PersonData struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	City    string `json:"city"`
}
type Location struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}
