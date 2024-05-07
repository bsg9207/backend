// Created by Seunggwan, Back on 2024.04.22
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package util

import (
	"fmt"
	"gorani/api/logHandler"
	"gorani/common/errors"

	"github.com/gin-gonic/gin"
)

func APIErrorHandler(_c *gin.Context, _requestId string) {
	s := recover()
	if s != nil {
		// convert to TError
		recovered, ok := s.(errors.TError)
		if !ok {
			err := fmt.Errorf("%v", s)
			recovered = errors.ERROR_UNDEFINED(err.Error()).(errors.TError)
		}

		logHandler.Write("trace", 0, recovered.Error())
		BadRequestResponse(
			_c,
			MakeErrorResponseString(_requestId, recovered, false),
		)
	}
}
