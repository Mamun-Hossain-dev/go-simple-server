package database

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

var users []User

// Store user
func (u User) Store() User {
	if u.ID != 0 {
		return u
	}

	u.ID = len(users) + 1
	users = append(users, u)
	return u
}

// Find user
func Find(email, pass string) *User {
	for idx, user := range users {
		if (user.Email == email) && (user.Password == pass) {
			return &users[idx]
		}
	}

	return nil
}
