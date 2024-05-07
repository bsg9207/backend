// Created by Seunggwan, Back on 2024.04.25
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package rpcRequest

import (
	"fmt"
	"testing"

	// testify assert
	"github.com/stretchr/testify/assert"
)

func Test_RpcRequest(t *testing.T) {
	// ready rpc request
	const rpcUrl = "http://192.168.0.201:8545"
	const method = "eth_getBlockByNumber"
	params := []interface{}{fmt.Sprintf("0x%x", -1), true}
	const id = 1

	// do rpc request
	res := RpcRequest(rpcUrl, method, params, id)

	assert.Equal(t, res.Result, nil, "result should be nil")
	assert.NotNil(t, res.Error)
}
