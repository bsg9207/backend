// Created by Seunggwan, Back on 2024.04.22
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package util

import (
	"gorani/common/errors"
	"testing"

	// testify assert
	"github.com/stretchr/testify/assert"
)

func Test_MakeResponse(t *testing.T) {
	type TTemp struct {
		ABC string `json:"abc"`
		DEF string `json:"def"`
	}

	res := MakeResponseString(
		"1",
		TTemp{
			"ABC",
			"DEF",
		},
		false,
	)

	assert.Equal(
		t,
		res,
		`{"requestId":"1","result":{"data":{"abc":"ABC","def":"DEF"}}}`,
	)
}

func Test_MakeErrorResponse(t *testing.T) {
	res := MakeErrorResponseString(
		"1",
		errors.ERROR_API_UNSUPPORTED_ROUTER_METHOD().(errors.TError),
		false,
	)

	assert.Equal(
		t,
		res,
		`{"requestId":"1","result":{"code":"20001","error":"Unsupported router method"}}`,
	)
}
