package mygoapp

type UserCredentials struct {
	Email string
	Password string
}

type User struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Phone   string   `json:"phone"`
	Address *Address `json:"address"`
}

type Address struct {
	Id      string `json:"id"`
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type LoginService interface {
	Authenticate(UserCredentials)
	Login(UserCredentials)
}

type Repository interface {
	Get(id string) User
}
