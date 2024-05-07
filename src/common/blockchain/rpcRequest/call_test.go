// Created by Seunggwan, Back on 2024.04.26
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package rpcRequest

import (
	"testing"

	// testify assert
	"github.com/stretchr/testify/assert"
)

func Test_Eth_Call(t *testing.T) {
	// ready rpc request
	const url = "https://kaikas.baobab.klaytn.net:8651"
	rpc := &TRPC{url}

	// ready params
	const toAddress = "0xf91043c37d89c7768d87f59b8eba0ac3b7d83b53"
	const fromAddress = "0x774cf653f3c0221dc5696721f0cd6fb52b4bb593"
	const data = "0xfa288b000000000000000000000000009e056611871794d0a4ece7d0dc014e34e7469d6a"
	const block = "latest"
	const id = 1
	res := rpc.Call(fromAddress, toAddress, data, block, id)

	assert.Equal(t, res.ResultToString(), "0x000000000000000000000000e9119ba33d4fff07cebf5a5f9f58a1ba14e127fd0000000000000000000000003cae9255b7ad17df118ed4f2338ff80c3ee3a15e")
	assert.Equal(t, res.Error, nil)
}
