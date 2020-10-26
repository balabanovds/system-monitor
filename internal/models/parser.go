package models

import (
	"bytes"
	"encoding/json"
)

type ParserType uint8

const (
	Undef ParserType = iota
	LoadAvg
	CPU
	IO
	FS
	Net
)

func (t ParserType) String() string {
	return toString[t]
}

func (t ParserType) Value(str string) ParserType {
	v, ok := toID[str]
	if !ok {
		return Undef
	}
	return v
}

var toString = map[ParserType]string{
	LoadAvg: "load_avg",
	CPU:     "cpu",
	IO:      "io",
	FS:      "fs",
	Net:     "net",
}

var toID = map[string]ParserType{
	"load_avg": LoadAvg,
	"cpu":      CPU,
	"io":       IO,
	"fs":       FS,
	"net":      Net,
}

func (t ParserType) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(t.String())
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (t *ParserType) UnmarshalJSON(b []byte) error {
	var j string
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}

	*t = toID[j]
	return nil
}
