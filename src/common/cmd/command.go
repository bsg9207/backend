// Created by Seunggwan, Back on 2024.04.23
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package cmd

import (
	"flag"
	"gorani/common/errors"
)

type TArgType string

const (
	STRING TArgType = "string"
	INT    TArgType = "int"
	BOOL   TArgType = "bool"
)

func GetArg(
	_type TArgType,
	_name string,
	_defaultValue interface{},
	_desc string,
) interface{} {
	switch _type {
	case STRING:
		value := _defaultValue.(string)
		arg := flag.String(_name, value, _desc)
		flag.Parse()
		return arg
	case INT:
		value := _defaultValue.(int)
		arg := flag.Int(_name, value, _desc)
		flag.Parse()
		return arg
	case BOOL:
		value := _defaultValue.(bool)
		arg := flag.Bool(_name, value, _desc)
		flag.Parse()
		return arg
	default:
		panic(errors.ERROR_CMD_UNSUPPORTED_ARGUMENT_TYPE())
	}
}
