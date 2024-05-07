// Created by Seunggwan, Back on 2024.04.17
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package errors

import "gorani/common/utils"

type TError struct {
	Code        string `json:"code"`  // error code
	CustomError string `json:"error"` // error message
}

func (e TError) String() string {
	return utils.InterfaceToJsonString(e, true)
}

func (e TError) Error() string {
	return e.String()
}
