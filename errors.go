package logger

import "errors"

var (
	AlreadyConnectedErr  = errors.New("Already connected to server")
	ServerConnectionFail = errors.New("Server Connection failed")
	TypeNotMatch         = errors.New("Type not match")
)
