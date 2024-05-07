// Created by Seunggwan, Back on 2024.04.22
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package section

// key value
// ${section_name}.${key_name}
const (
	// section name
	SECTION_SERVER = "server"

	// string
	KEY_SERVER_HOST = SECTION_SERVER + ".host"

	// bool
	KEY_SERVER_DEBUG = SECTION_SERVER + ".debug"
)

// //////////////////////////////////////////////////////////////////////////////
// section - server
type TServer struct {
	Host  string
	Debug bool
}
