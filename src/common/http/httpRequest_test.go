// Created by Seunggwan, Back on 2024.04.24
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package http

import (
	"gorani/common/utils"
	"testing"

	// testify assert
	"github.com/stretchr/testify/assert"
)

func Test_GetRequestToBytes(t *testing.T) {
	url := "http://192.168.0.201:8090/management/health/ping"
	timeout := 5 * 1000

	result, _ := GetRequestToBytes(url, timeout)
	assert.Equal(
		t,
		utils.BytesToString(result),
		`{"requestId":"0","result":{"data":"pong"}}`,
	)
}

func Test_GetRequestToString(t *testing.T) {
	url := "http://192.168.0.201:8090/management/health/ping"
	timeout := 5 * 1000

	result, _ := GetRequestToString(url, timeout)
	assert.Equal(
		t,
		result,
		`{"requestId":"0","result":{"data":"pong"}}`,
	)
}

func Test_GetRequestToJson(t *testing.T) {
	url := "http://192.168.0.201:8090/management/health/ping"
	timeout := 5 * 1000

	result, _ := GetRequestToJson(url, timeout)
	assert.Equal(
		t,
		result.PPrint(),
		`{
  "requestId": "0",
  "result": {
    "data": "pong"
  }
}`,
	)
}
