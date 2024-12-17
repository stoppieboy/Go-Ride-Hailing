package User

type User struct {
	id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Ride     []Ride `json:"ride"`
}

type Ride struct {
	id          string `json:"id"`
	Source      string `json:"source"`
	Destination string `json:"dest"`
	Rider       string `json:"user"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Success bool   `json:"success"`
}
