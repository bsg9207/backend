// Created by Seunggwan, Back on 2024.04.22
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package routers

// default table
func DefaultAPITable() TRouterTable {
	rt := TRouterTable{}

	// ping
	rt.AddGetRouter("/management/health/ping", Ping, false)

	return rt
}
