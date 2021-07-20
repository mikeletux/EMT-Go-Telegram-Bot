package auth

// Auth is the interface structs need to comply with to be able to be used
// for autentication within the bot
type Auth interface {
	// CheckUser should return if the user is allows to user the service
	CheckUser(username string) error
	Register(username string) error
}
