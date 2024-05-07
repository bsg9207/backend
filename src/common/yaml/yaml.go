// Created by Seunggwan, Back on 2024.04.18
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package yaml

import (
	"fmt"
	"io/ioutil"

	"gorani/common/errors"

	"gopkg.in/yaml.v2"
)

func LoadFromFile(_strFilePath string, _mapUnmarshal *map[string]interface{}) error {
	byteData, err := ioutil.ReadFile(_strFilePath)
	if err != nil {
		return errors.ERROR_YAML_READ_FILE(err.Error())
	}
	return LoadFromBytes(byteData, _mapUnmarshal)
}

func LoadFromBytes(_byteData []byte, _mapUnmarshal *map[string]interface{}) error {
	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal(_byteData, &m)
	if err != nil {
		return errors.ERROR_YAML_UNMARSHAL(err.Error())
	}

	for k := range m {
		strNewKey := fmt.Sprintf("%v", k)
		(*_mapUnmarshal)[strNewKey] = m[k]
		err = _marshaling(strNewKey, m[k], _mapUnmarshal)
		if err != nil {
			break
		}
	}

	m = nil
	return nil
}

func _marshaling(_strKey string, _m interface{}, _mapUnmarshal *map[string]interface{}) error {
	byteData, err := yaml.Marshal(_m)
	if err != nil {
		return errors.ERROR_YAML_MARSHAL(err.Error())
	}

	unMar := make(map[interface{}]interface{})
	err = yaml.Unmarshal(byteData, &unMar)
	if err != nil {
		return errors.ERROR_YAML_UNMARSHAL(err.Error())
	}

	for k := range unMar {
		strCurrentKey := fmt.Sprintf("%v", k)
		strNewKey := _strKey + "." + strCurrentKey
		(*_mapUnmarshal)[strNewKey] = unMar[k]
		_marshaling(strNewKey, unMar[k], _mapUnmarshal)
	}
	unMar = nil
	byteData = nil

	return nil
}
