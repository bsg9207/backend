// Created by Seunggwan, Back on 2024.04.17
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package utils

import (
	"strconv"

	"gorani/common/utils/json"
)

// bytes to string
func BytesToString(_bytes []byte) string {
	return string(_bytes)
}

// I64ToString : Int64 to String
func I64ToString(_i64 int64, _base int) string {
	return strconv.FormatInt(_i64, _base)
}

// StringToI64 : String to Int64
func StringToI64(_str string, _base int) int64 {
	i, err := strconv.ParseInt(_str, _base, 64)
	if err != nil {
		panic(err)
	}
	return i
}

// StringToUin64 : String to Uint64
func StringToUint64(_str string, _base int) uint64 {
	i, err := strconv.ParseUint(_str, _base, 64)
	if err != nil {
		panic(err)
	}
	return i
}

// interface to json string
func InterfaceToJsonString(_interface interface{}, _pretty bool) string {
	bytes, err := json.ToJson(_interface)
	if err != nil {
		panic(err)
	}

	jMap, err := json.NewBytes(bytes)
	if err != nil {
		panic(err)
	}

	if _pretty {
		return jMap.PPrint()
	} else {
		return jMap.Print()
	}
}

// interface array to string array
func InterfaceArrayToStringArray(_interface []interface{}) []string {
	if _interface != nil {
		result := []string{}
		for _, inter := range _interface {
			result = append(result, inter.(string))
		}

		return result
	}

	return []string{}
}
