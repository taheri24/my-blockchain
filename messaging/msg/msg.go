package msg

// this file for `Msg` data structure , it contains `Msg struct`  prototype(fields,constructor) Only

import "encoding/json"

type Msg struct {
	SenderID  string          `json:"sender"`
	Signature string          `json:"signature"`
	Contents  json.RawMessage `json:"contents"`
}

func New(VisitFns ...VisitFunc) *Msg {
	m := new(Msg)
	for _, visit := range VisitFns {
		visit(m)
	}

	return m
}
