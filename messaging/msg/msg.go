package msg

// this file for `Msg` data structure , it contains `Msg struct`  prototype(fields,constructor) Only

import "encoding/json"

type Msg struct {
	SenderID       string          `json:"sender"`
	Signature      string          `json:"signature"`
	Contents       json.RawMessage `json:"contents"`
	ContentsFormat int8            `json:"format"`
	ContentsKind   int8            `json:"kind"`
}

func New(opts ...Option) *Msg {
	m := new(Msg)
	applyOptions(m, opts)

	return m
}
