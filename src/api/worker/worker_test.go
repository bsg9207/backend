// Created by Seunggwan, Back on 2024.04.22
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package worker

import (
	"gorani/api/preference"
	"gorani/api/routers"
	"testing"
	// testify assert
)

func Test_run(t *testing.T) {
	preference.Initialize("local")

	Run(routers.DefaultAPITable())
}
