package model

const (
	AUTH_FLAG_LOGGED_IN = iota
	AUTH_FLAG_LOGGED_OUT
)

type Auth struct {
	Id     int32
	IdUser int32
	Token  string
	Flags  int32
}
