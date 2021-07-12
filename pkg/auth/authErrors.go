package auth

type UserNotFoundError string

func (u UserNotFoundError) Error() string {
	return string(u)
}
