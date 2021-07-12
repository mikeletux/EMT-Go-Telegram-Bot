package auth

import "fmt"

type simpleAuth struct {
	AllowedUsers []string
}

func NewSimpleAuth(AllowedUsers []string) Auth {
	simpleAuth := simpleAuth{
		AllowedUsers: AllowedUsers,
	}
	return &simpleAuth
}

func (s *simpleAuth) CheckUser(username string) error {
	for _, v := range s.AllowedUsers {
		if v == username {
			return nil
		}
	}
	return UserNotFoundError(fmt.Sprintf("the user %s was not found in the system", username))
}
