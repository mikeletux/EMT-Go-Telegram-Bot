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

func (s *simpleAuth) Register(username string) error {
	if len(username) == 0 {
		return InternalError("username cannot be empty")
	}
	// Check that the user does not exist already
	for _, v := range s.AllowedUsers {
		if v == username {
			return UserAlreadyExists(fmt.Sprintf("the user %s already exists", username))
		}
	}

	s.AllowedUsers = append(s.AllowedUsers, username)

	return nil
}
