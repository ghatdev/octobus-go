package octobus

import (
	"encoding/json"
	"net"
	"time"
)

type BasicLogger struct {
	Config *Config
	conn   net.Conn
}

// Dial to Server
// Now for using TCP or HTTP
func (l *BasicLogger) Dial() error {
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
func (l *BasicLogger) Log(t, tag string, value interface{}) error {
	logPayload := log{Type: t, Tag: tag, Service: l.Config.Service, Time: time.Now()}

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
