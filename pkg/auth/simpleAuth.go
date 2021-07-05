package auth

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
