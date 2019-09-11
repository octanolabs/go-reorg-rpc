package main

import (
	"github.com/octanolabs/go-reorg-rpc/common"
)

type request struct {
	Jsonrpc string   `json:"jsonrpc"`
	Id      int      `json:"id"`
	Method  string   `json:"method"`
	Params  []string `json:"params"` // TODO, json.RawMessage
}

type strPayload struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  string `json:"result"`
}

type blockPayload struct {
	Jsonrpc string       `json:"jsonrpc"`
	Id      int          `json:"id"`
	Result  common.Block `json:"result"`
}
