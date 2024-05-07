// Created by Seunggwan, Back on 2024.04.25
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package rpcRequest

import (
	"encoding/json"
	"fmt"

	"gorani/common/errors"
	"gorani/common/utils"
	jMapper "gorani/common/utils/json"
)

type TPayload struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int64         `json:"id"`
}

func (p TPayload) ToBytes() []byte {
	bytes, err := json.Marshal(p)
	if err != nil {
		panic(errors.ERROR_RPC_PAYLOAD_TO_BYTES(err.Error()))
	}

	return bytes
}

// define rpc response structure
type TRpcResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int64       `json:"id"`
	Result  interface{} `json:"result"`
	Error   interface{} `json:"error"`
}

func (res *TRpcResponse) ToString(_pretty bool) string {
	bytesResult, err := jMapper.ToJson(res)
	if err != nil {
		panic(errors.ERROR_RPC_RESPONSE_TO_STRING(err.Error()))
	}

	jMap, err := jMapper.NewBytes(bytesResult)
	if err != nil {
		panic(errors.ERROR_RPC_RESPONSE_TO_STRING(err.Error()))
	}

	if _pretty {
		return jMap.PPrint()
	} else {
		return jMap.Print()
	}
}

// result to string
func (res *TRpcResponse) ResultToString() string {
	if res.Result != nil {
		return fmt.Sprintf("%v", res.Result)
	} else {
		return ""
	}
}

// result to json string
func (res *TRpcResponse) ResultToJsonString(pretty bool) string {
	if res.Result != nil {
		return utils.InterfaceToJsonString(res.Result, pretty)
	} else {
		return ""
	}
}

// result to json bytes
func (res *TRpcResponse) ResultToJson() *jMapper.TJsonMap {
	bytesError, err := jMapper.ToJson(res.Result)
	if err != nil {
		panic(errors.ERROR_RPC_RESPONSE_RESULT_TO_JSON(err.Error()))
	}

	jMap, err := jMapper.NewBytes(bytesError)
	if err != nil {
		panic(errors.ERROR_RPC_RESPONSE_RESULT_TO_JSON(err.Error()))
	}

	return jMap
}

// error to string
func (res *TRpcResponse) ErrorToString(pretty bool) string {
	jMap := res.ErrorToJson()

	if pretty {
		return jMap.PPrint()
	} else {
		return jMap.Print()
	}
}

// error to json bytes
func (res *TRpcResponse) ErrorToJson() *jMapper.TJsonMap {
	bytesError, err := jMapper.ToJson(res.Error)
	if err != nil {
		panic(errors.ERROR_RPC_RESPONSE_ERROR_TO_JSON(err.Error()))
	}

	jMap, err := jMapper.NewBytes(bytesError)
	if err != nil {
		panic(errors.ERROR_RPC_RESPONSE_ERROR_TO_JSON(err.Error()))
	}

	return jMap
}
