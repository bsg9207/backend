// Created by Seunggwan, Back on 2024.04.17
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package utils

import (
	"testing"

	// testify assert
	"github.com/stretchr/testify/assert"
)

func Test_InterfaceToJsonString(t *testing.T) {
	type TBindAddressResponse struct {
		UserId    string `json:"uid"`
		RequestId string `json:"requestId"`
		Success   bool   `json:"success"`
	}

	result := InterfaceToJsonString(
		TBindAddressResponse{
			"uid",
			"requestId",
			true,
		},
		false,
	)

	assert.Equal(
		t,
		result,
		`{"requestId":"requestId","success":true,"uid":"uid"}`,
	)
}
