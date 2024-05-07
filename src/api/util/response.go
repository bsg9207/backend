// Created by Seunggwan, Back on 2024.04.22
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package util

import (
	"gorani/common/errors"
	"gorani/common/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TAPIResult struct {
	// json 형태의 struct까지 지원
	Data interface{} `json:"data"`
}

type TAPIResponse struct {
	Id     string     `json:"requestId"`
	Result TAPIResult `json:"result"`
}

func (r TAPIResponse) ToString(_pretty bool) string {
	return utils.InterfaceToJsonString(r, _pretty)
}

func (r TAPIResponse) String() string {
	return r.ToString(false)
}

type TAPIErrorResponse struct {
	Id     string        `json:"requestId"`
	Result errors.TError `json:"result"`
}

func (r TAPIErrorResponse) ToString(_pretty bool) string {
	return utils.InterfaceToJsonString(r, _pretty)
}

func (r TAPIErrorResponse) String() string {
	return r.ToString(false)
}

func MakeResponseString(
	_requestId string,
	_data interface{},
	_pretty bool,
) string {
	res := TAPIResponse{
		_requestId,
		TAPIResult{_data},
	}

	return res.ToString(_pretty)
}

func MakeErrorResponseString(
	_requestId string,
	_error errors.TError,
	_pretty bool,
) string {
	res := TAPIErrorResponse{
		_requestId,
		_error,
	}

	return res.ToString(_pretty)
}

// 200 OK
func OKResponse(
	_c *gin.Context,
	_response string,
) {
	// send response
	sendResponse(_c, http.StatusOK, _response)
}

// 400 Bad Request
func BadRequestResponse(
	_c *gin.Context,
	_response string,
) {
	sendResponse(_c, http.StatusBadRequest, _response)
}

// 401 Unauthorized
func UnauthorizedResponse(
	_c *gin.Context,
	_response string,
) {
	// send response
	sendResponse(_c, http.StatusUnauthorized, _response)
}

func sendResponse(
	_c *gin.Context,
	_httpCode int,
	_response string,
) {
	// header
	SetDefaultHeader(_c)

	// response
	_c.String(_httpCode, _response)
}
