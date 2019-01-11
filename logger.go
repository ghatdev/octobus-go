package octobus

import (
	"time"
)

type Octobus struct {
	Logger
}

// Logger is interface for using this logger
//
type Logger interface {
	Log(string, string, interface{}) error
	Dial() error
}

type log struct {
	Service string
	Type    string
	Tag     string
	Value   string
	Time    time.Time `json:"_id" bson:"_id"`
}

type M map[interface{}]interface{}

// New func
// Make new default logger
func New(logger Logger) *Octobus {
	return &Octobus{Logger: logger}
}
