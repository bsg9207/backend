// Created by Seunggwan, Back on 2024.04.17
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package json

import "encoding/json"

// NewBytes : new jmap from json bytes
func NewBytes(_b []byte) (*TJsonMap, error) {
	j := &TJsonMap{
		m: make(map[string]interface{}),
	}
	err := FromJson(_b, &j.m)
	return j, err
}

// ToJson : object(struct) to json bytes
func ToJson(_o interface{}) ([]byte, error) {
	jsonBytes, err := json.Marshal(_o)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

// FromJson : json bytes to object(struct)
func FromJson(_byte []byte, _o interface{}) error {
	err := json.Unmarshal(_byte, &_o)
	return err
}
