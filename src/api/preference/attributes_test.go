// Created by Seunggwan, Back on 2024.04.22
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package preference

import (
	"testing"

	// testify assert
	"github.com/stretchr/testify/assert"
)

func Test_GetAttributes(t *testing.T) {
	// initialize preference
	Initialize("local")

	// section - server
	host := GetHost()
	assert.Equal(
		t,
		host,
		"0.0.0.0:8080",
	)

	isDebug := IsDebugMode()
	assert.Equal(
		t,
		isDebug,
		true,
	)
}
