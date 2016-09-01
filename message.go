package gochill

/*
This library should be take responsible just to build log data structure.

Expected usages :

  log.Println(gochill.Info("short message"))

Each of method should return json string with valid structured data.
Each of method should initiate new instance of Message struct, we need to
do this to initiate host, service and timestamp value.
*/

import (
	"encoding/json"
	"os"
	"time"
)

//Message struct used to build all required fields
type Message struct {
	Version      int    `json:"version"`
	Host         string `json:"host"`
	Service      string `json:"service"`
	ShortMessage string `json:"short_message"`
	Timestamp    int64  `json:"timestamp"`
	Level        int    `json:"level"`
	jsonEngine   JSONMarshalEngine
	osEngine     OSHostEngine
}

//ToJSON used to convert Message struct from byte to string after marshalled by
//json encoding
func (m *Message) ToJSON() []byte {
	jsonByte, err := m.jsonEngine(m)
	if err != nil {
		return nil
	}

	return jsonByte
}

//ReplaceJSONEngine used to change json engine
func (m *Message) ReplaceJSONEngine(je JSONMarshalEngine) {
	m.jsonEngine = je
}

//ReplaceOSEngine used to change native golang os hostname implementation
func (m *Message) ReplaceOSEngine(oe OSHostEngine) {
	m.osEngine = oe
	m.parseHostName()
}

//parseHostName used to parse os hostname, but to simplify our library
//i just create this function, so we doesn't need to care about multiple
//return value for caller, is something goes wrong just return an empty string
func (m *Message) parseHostName() {
	host, err := m.osEngine()
	if err != nil {
		m.Host = ""
	}

	m.Host = host
}

//NewMessage used to create new instance of `message` struct
func NewMessage(level int) *Message {
	m := Message{}
	m.Version = Version
	m.Service = os.Getenv(EnvServiceKeyName)
	m.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	m.Level = level
	m.jsonEngine = json.Marshal
	m.osEngine = os.Hostname
	m.parseHostName()
	return &m
}
