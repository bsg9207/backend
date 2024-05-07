// Created by Seunggwan, Back on 2024.04.17
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package json

import (
	"bytes"
	"encoding/json"
	"reflect"
)

type TJsonMap struct {
	m map[string]interface{}

	// for finder/remover
	splitKey []string
	cursor   int

	// for adder
	insertKey   string
	insertValue interface{}
}

/*** reflect type definition ***/
var IntType = reflect.TypeOf(1)
var Float64Type = reflect.TypeOf(float64(1))
var StringType = reflect.TypeOf(string(""))
var BoolType = reflect.TypeOf(false)
var SliceType = reflect.TypeOf([]interface{}(nil))
var SliceMapType = reflect.TypeOf([]map[string]interface{}(nil))
var JsonMapType = reflect.TypeOf((map[string]interface{})(nil))

/*** Token for key depth ***/
var SPLIT_TOKEN string = "."

// PPrint : pretty print
func (j *TJsonMap) PPrint() string {
	b, _ := ToJson(j.m)
	str, _ := prettyPrint(b)
	return str
}

// prettyPrint : pretty-print byte to json string
func prettyPrint(_src []byte) (string, error) {
	var dst bytes.Buffer
	err := json.Indent(&dst, _src, "", "  ")
	if nil != err {
		return "", err
	}
	return dst.String(), nil
}

// Print : print
func (j *TJsonMap) Print() string {
	b, _ := ToJson(j.m)
	return string(b)
}
