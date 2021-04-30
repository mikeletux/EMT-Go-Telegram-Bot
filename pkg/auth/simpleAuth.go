package auth

type Auth interface {
	// CheckUser should return if the user is allows to user the service
	CheckUser(username string) bool
}

type simpleAuth struct {
	AllowedUsers []string
}

func NewSimpleAuth(AllowedUsers []string) Auth {
	simpleAuth := simpleAuth{
		AllowedUsers: AllowedUsers,
	}
	return &simpleAuth
}

func (s *simpleAuth) CheckUser(username string) bool {
	for _, v := range s.AllowedUsers {
		if v == username {
			return true
		}
	}
	return false
}
