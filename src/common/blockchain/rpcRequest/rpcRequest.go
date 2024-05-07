// Created by Seunggwan, Back on 2024.04.25
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package rpcRequest

import (
	"bytes"
	"encoding/json"

	"gorani/common/errors"
	"gorani/common/http"
)

func RpcRequest(_url string, _method string, _params []interface{}, _id int64) *TRpcResponse {
	// ready body
	data := &TPayload{
		Jsonrpc: "2.0",
		Method:  _method,
		Params:  _params,
		ID:      _id,
	}

	// do http request
	// timeout : 30 sec
	const timeout = 30 * 1000
	jsonResponse, err := http.PostRequestFromBytes(
		_url, data.ToBytes(), timeout)
	if err != nil {
		panic(errors.ERROR_RPC_REQUEST(err.Error()))
	}

	// convert to bytes from response
	bytes := bytes.NewBufferString(jsonResponse.PPrint())
	if err != nil {
		panic(errors.ERROR_RPC_REQUEST(err.Error()))
	}

	// result from bytes
	result := &TRpcResponse{}
	err = json.Unmarshal(bytes.Bytes(), result)
	if err != nil {
		panic(errors.ERROR_RPC_REQUEST(err.Error()))
	}

	data = nil
	jsonResponse = nil
	bytes = nil

	// release memory
	defer func() {
		result = nil
	}()

	return result
}
