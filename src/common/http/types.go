// Created by Seunggwan, Back on 2024.04.25
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package http

type THttpMethod string

func (tm THttpMethod) String() string {
	return string(tm)
}

const (
	GET  THttpMethod = "GET"
	POST THttpMethod = "POST"
)
