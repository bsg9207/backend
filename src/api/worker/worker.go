// Created by Seunggwan, Back on 2024.04.22
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package worker

import (
	"gorani/api/preference"
	"gorani/api/routers"
	"gorani/common/errors"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(_router routers.TRouterTable) error {
	return worker(_router)
}

func worker(_router routers.TRouterTable) error {
	// gin mode
	var r *gin.Engine
	if preference.IsDebugMode() {
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// set middleware
	onSetMiddleware(r)

	// set router
	onSetRouters(r, _router)

	// start gin
	r.Run(preference.GetHost())

	return nil
}

func onSetMiddleware(_r *gin.Engine) {
	if !preference.IsDebugMode() {
		_r.Use(gin.Recovery())
	}

	_r.Use(
		cors.New(
			cors.Config{
				AllowOrigins: []string{"*"},
				AllowMethods: []string{
					routers.GET.String(),
					routers.POST.String(),
				},
				AllowHeaders:     []string{"*"},
				AllowCredentials: true,
			},
		),
	)
}

func onSetRouters(_r *gin.Engine, _router routers.TRouterTable) {
	// TODO : path prefix
	authPath := "/"
	authRequired := _r.Group(authPath)

	// TODO : auth middleware

	// no auth
	noAuth := _r.Group("/")

	for _, props := range _router {
		if props.Auth {
			_setRouter(authRequired, props)
		} else {
			_setRouter(noAuth, props)
		}
	}
}

func _setRouter(_rg *gin.RouterGroup, _props routers.TRouterProps) {
	switch _props.Method {
	case routers.GET:
		_rg.GET(_props.Path, _props.Function)
	case routers.POST:
		_rg.POST(_props.Path, _props.Function)
	default:
		panic(errors.ERROR_API_UNSUPPORTED_ROUTER_METHOD())
	}
}
