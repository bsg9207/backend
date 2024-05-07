// Created by Seunggwan, Back on 2024.04.22
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package preference

import (
	"fmt"
	"gorani/common/utils"
	"math/big"
)

func _find(_key string) interface{} {
	inst := GetInstance()

	// lock
	inst.lockMutex.Lock()
	defer inst.lockMutex.Unlock()

	return inst.mapYaml[_key]
}

func _findToString(_key string, _out *string) {
	*_out = fmt.Sprint(_find(_key))
	if *_out == "<nil>" {
		*_out = ""
	}
}

func _findToInt(_key string, _out *int) {
	// find to int64
	nVal := int64(0)
	_findToInt64(_key, &nVal)

	// assign to out
	*_out = int(nVal)
}

func _findToInt64(_key string, _out *int64) {
	// get from yaml
	strVal := ""
	_findToString(_key, &strVal)

	// check value
	if strVal == "" {
		*_out = int64(0)
		return
	}

	// string to int64
	*_out = utils.StringToI64(strVal, 10)
}

func _findToUint64(_key string, _out *uint64) {
	// get from yaml
	strVal := ""
	_findToString(_key, &strVal)

	// check value
	if strVal == "" {
		*_out = uint64(0)
		return
	}

	// string to uint64
	*_out = utils.StringToUint64(strVal, 10)
}

func _findToBigInt(_key string, _out **big.Int) {
	// get from yaml (int64)
	i64 := int64(0)
	_findToInt64(_key, &i64)

	// int64 to big int
	*_out = big.NewInt(i64)
}

func _findToBool(_key string, _out *bool) {
	// get from yaml (interface{})
	flag := _find(_key)

	if flag == nil {
		*_out = false
	}

	// interface{} to bool
	*_out = flag.(bool)
}

func _findToStringArray(_key string, _out *[]string) {
	// get from yaml (interface{})
	arrValue, ok := _find(_key).([]interface{})
	if ok {
		*_out = utils.InterfaceArrayToStringArray(arrValue)
	}
}
