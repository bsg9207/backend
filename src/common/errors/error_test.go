// Created by Seunggwan, Back on 2024.04.17
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package errors

import (
	"testing"

	// testify assert
	"github.com/stretchr/testify/assert"
)

func Test_CheckIsError(t *testing.T) {
	// create custom error
	customError := TError{}

	// convert to error
	_, ok := interface{}(customError).(error)

	assert.Equal(
		t,
		ok,
		true,
	)
}

func Test_ConvertFromError(t *testing.T) {
	defer func() {
		s := recover()
		if s != nil {
			recoveredCustomError := s.(TError)

			LogError(recoveredCustomError)
			assert.Equal(
				t,
				recoveredCustomError.String(),
				"Test Message",
			)
		}
	}()

	panic(TError{"1000", "Test Message"})
}

func Test_ERROR_TEST(t *testing.T) {
	errTest := ERROR_TEST().(TError)
	assert.Equal(
		t,
		errTest.Code,
		"-1",
	)
	assert.Equal(
		t,
		errTest.CustomError,
		"Test Message",
	)
	assert.Equal(
		t,
		errTest.Error(),
		"Test Message",
	)
}
