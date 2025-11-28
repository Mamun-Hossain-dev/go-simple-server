package user

type UserRepository interface {
	Store(u User) User
	Find(email, pass string) *User
}

type UserRepo struct {
}

func NewUserRepository() UserRepository {
	return &UserRepo{}
}

func (u *UserRepo) Store(user User) User {
	if user.ID != 0 {
		return user
	}

	user.ID = len(Users) + 1
	Users = append(Users, user)
	return user
}

func (u *UserRepo) Find(email, pass string) *User {
	for idx, user := range Users {
		if (user.Email == email) && (user.Password == pass) {
			return &Users[idx]
		}
	}
	return nil
}
