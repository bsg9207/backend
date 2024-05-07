// Created by Seunggwan, Back on 2024.04.26
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package rpcRequest

type TRPCMethod string

const (
	ETH_CALL TRPCMethod = "eth_call"
)

func (m TRPCMethod) String() string {
	return string(m)
}

type TRPC struct {
	URL string // rpc endpoint url
}

func (trpc *TRPC) Request(_method string, _params []interface{}, _id int64) *TRpcResponse {
	result := RpcRequest(trpc.URL, _method, _params, _id)

	// release memory
	defer func() {
		result = nil
	}()

	return result
}
