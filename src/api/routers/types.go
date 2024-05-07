// Created by Seunggwan, Back on 2024.04.22
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package routers

import (
	"gorani/common/http"

	"github.com/gin-gonic/gin"
)

type TRouterProps struct {
	Method   http.THttpMethod
	Path     string
	Function gin.HandlerFunc
	Auth     bool
}

func genRouterProps(
	_method http.THttpMethod,
	_path string,
	_func gin.HandlerFunc,
	_auth bool,
) TRouterProps {
	return TRouterProps{
		Method:   _method,
		Path:     _path,
		Function: _func,
		Auth:     _auth,
	}
}

func GenGetRouterProps(
	_path string,
	_func gin.HandlerFunc,
	_auth bool,
) TRouterProps {
	return genRouterProps(http.GET, _path, _func, _auth)
}

func GenPostRouterProps(
	_path string,
	_func gin.HandlerFunc,
	_auth bool,
) TRouterProps {
	return genRouterProps(http.POST, _path, _func, _auth)
}

// router table
type TRouterTable []TRouterProps

func (rt *TRouterTable) AddRouter(
	_props TRouterProps,
) {
	*rt = append(*rt, _props)
}

func (rt *TRouterTable) AddGetRouter(
	_path string,
	_func gin.HandlerFunc,
	_auth bool,
) {
	rt.AddRouter(
		GenGetRouterProps(
			_path,
			_func,
			_auth,
		),
	)
}

func (rt *TRouterTable) AddPostRouter(
	_path string,
	_func gin.HandlerFunc,
	_auth bool,
) {
	rt.AddRouter(
		GenPostRouterProps(
			_path,
			_func,
			_auth,
		),
	)
}
