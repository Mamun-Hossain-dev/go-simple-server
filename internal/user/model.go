package user

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"isAdmin"`
}

type LoggedUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRes struct {
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type LoginResponse struct {
	Message      string `json:"message"`
	User         User   `json:"user"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
