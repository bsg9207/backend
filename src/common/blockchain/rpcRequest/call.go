// Created by Seunggwan, Back on 2024.04.26
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package rpcRequest

type TCallData struct {
	From string `json:"from"`
	To   string `json:"to"`
	Data string `json:"data"`
}

func (rpc *TRPC) Call(
	_from,
	_to,
	_data,
	_block string,
	_id int64,
) *TRpcResponse {
	return rpc.Request(
		ETH_CALL.String(),
		[]interface{}{
			&TCallData{_from, _to, _data},
			_block,
		},
		_id,
	)
}
