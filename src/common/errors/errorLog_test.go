// Created by Seunggwan, Back on 2024.04.18
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package errors

import (
	"testing"

	// testify assert
	"github.com/stretchr/testify/assert"
)

func Test_LogError(t *testing.T) {
	errTest := ERROR_TEST()

	assert.Equal(
		t,
		LogError(errTest),
		true,
	)
}
