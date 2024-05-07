// Created by Seunggwan, Back on 2024.04.22
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package util

import "github.com/gin-gonic/gin"

func GetIdFromHeader(_c *gin.Context) string {
	// set id
	var id string
	reqId := _c.GetHeader("ID")
	if reqId != "" {
		id = reqId
	} else {
		id = "0"
	}

	return id
}

func SetDefaultHeader(_c *gin.Context) {
	_c.Header("Content-Type", "application/json; charset=utf-8")
}
