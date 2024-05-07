// Created by Seunggwan, Back on 2024.04.22
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package routers

import (
	"gorani/api/util"

	"github.com/gin-gonic/gin"
)

func Ping(_c *gin.Context) {
	// get id from header
	requestId := util.GetIdFromHeader(_c)

	// error handling
	defer util.APIErrorHandler(_c, requestId)

	// pong
	res := util.MakeResponseString(
		requestId,
		"pong",
		false,
	)

	// response
	util.OKResponse(_c, res)
}
