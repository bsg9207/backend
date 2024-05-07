// Created by Seunggwan, Back on 2024.04.19
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package logHandler

import (
	"testing"
)

func Test_Signer(t *testing.T) {

	Initialize("local")

	Write("trace", 0, "start\t", "key\t")

}
