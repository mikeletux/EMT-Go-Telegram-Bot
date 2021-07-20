package auth

type UserNotFoundError string

func (u UserNotFoundError) Error() string {
	return string(u)
}

type UserAlreadyExists string

func (u UserAlreadyExists) Error() string {
	return string(u)
}

type InternalError string

func (i InternalError) Error() string {
	return string(i)
}
