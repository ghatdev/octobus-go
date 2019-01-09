package logger

import (
	"encoding/json"
	"net"
	"time"
)

type Log interface {
	Log()
}

// Logger struct
type Logger struct {
	Config *Config
	conn   net.Conn
}

type log struct {
	Key     string
	Service string
	Type    string
	Tag     string
	Value   string
	Time    time.Time `json:"_id" bson:"_id"`
}

type M map[interface{}]interface{}

// New func
// Make new default logger
func New(c *Config) *Logger {
	return &Logger{Config: c}
}

// Dial to Server
// Now for using TCP or HTTP
func (l *Logger) Dial() error {
	conn, err := net.Dial("tcp", l.Config.Host)
	if err != nil {
		return err
	}

	l.conn = conn
	return nil
}

// Log func
// Logging a Log
// Marshals values as map
func (l *Logger) Log(key, t string, value interface{}) error {
	logPayload := log{Key: key, Type: t}

	switch value.(type) {
	case string: // case json-string or else plain string
		logPayload.Value = value.(string)
	case M: // case marshall map. Need to marshall to json or else (now for json)
		data, err := json.Marshal(value.(M))
		if err != nil {
			return err
		}

		logPayload.Value = string(data)
	default:
		return TypeNotMatch
	}

	return json.NewEncoder(l.conn).Encode(&logPayload)
}
