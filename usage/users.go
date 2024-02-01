package usage

import cache "genericCache"

type User struct {
	ID    string
	Name  string
	Email string
}

type UsersService struct {
	cache *cache.Cache[User]
}

func NewUsersService() *UsersService {
	return &UsersService{
		cache: cache.NewCache[User](),
	}
}

func (s *UsersService) AddUser(user User) {
	s.cache.Set(user.ID, user)
}

func (s *UsersService) GetUser(id string) User {
	return s.cache.Get(id)
}
