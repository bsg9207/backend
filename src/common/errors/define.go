// Created by Seunggwan, Back on 2024.04.17
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package errors

// //////////////////////////////////////////////////////////////////////////////
// common
// define code & error (code have to pair with error)
const (
	// test error
	TEST_CODE  = "0"
	TEST_ERROR = "Test Message"

	// undefined error
	ERROR_UNDEFINED_CODE = "-1"

	// common error
	// yaml
	YAML_READ_FILE_CODE  = "10001"
	YAML_MARSHAL_CODE    = "10002"
	YAML_UNMARSHAL_CODE  = "10003"
	YAML_MARSHALING_CODE = "10004"
	YAML_LOAD_FROM_BYTES = "10005"

	// file
	FILE_MKDIR = "11001"

	// cmd
	CMD_UNSUPPORTED_ARGUMENT_TYPE     = "12001"
	MSG_CMD_UNSUPPORTED_ARGUMENT_TYPE = "Unsupported argument type"

	// http
	HTTP_NEW_REQUEST  = "13001"
	HTTP_SEND_REQUEST = "13002"

	// blockchain
	// rpc request
	RPC_PAYLOAD_TO_BYTES        = "14001"
	RPC_RESPONSE_TO_STRING      = "14002"
	RPC_RESPONSE_RESULT_TO_JSON = "14003"
	RPC_RESPONSE_ERROR_TO_JSON  = "14004"
	RPC_REQUEST                 = "14005"

	// api error
	API_UNSUPPORTED_ROUTER_METHOD     = "20001"
	MSG_API_UNSUPPORTED_ROUTER_METHOD = "Unsupported router method"
)

// define custom type error
func ERROR_TEST() error {
	return TError{
		TEST_CODE,
		TEST_ERROR,
	}
}

// undefined error
func ERROR_UNDEFINED(_msg string) error { return TError{ERROR_UNDEFINED_CODE, _msg} }

// yaml
func ERROR_YAML_READ_FILE(_msg string) error       { return TError{YAML_READ_FILE_CODE, _msg} }
func ERROR_YAML_MARSHAL(_msg string) error         { return TError{YAML_MARSHAL_CODE, _msg} }
func ERROR_YAML_UNMARSHAL(_msg string) error       { return TError{YAML_UNMARSHAL_CODE, _msg} }
func ERROR_YAML_MARSHALING(_msg string) error      { return TError{YAML_MARSHALING_CODE, _msg} }
func ERROR_YAML_LOAD_FROM_BYTES(_msg string) error { return TError{YAML_LOAD_FROM_BYTES, _msg} }

// file
func ERROR_MKDIR(_msg string) error { return TError{FILE_MKDIR, _msg} }

// cmd
func ERROR_CMD_UNSUPPORTED_ARGUMENT_TYPE() error {
	return TError{CMD_UNSUPPORTED_ARGUMENT_TYPE, MSG_CMD_UNSUPPORTED_ARGUMENT_TYPE}
}

// http
func ERROR_HTTP_NEW_REQUEST(_msg string) error  { return TError{HTTP_NEW_REQUEST, _msg} }
func ERROR_HTTP_SEND_REQUEST(_msg string) error { return TError{HTTP_SEND_REQUEST, _msg} }

// blockchain
// rpc request
func ERROR_RPC_PAYLOAD_TO_BYTES(_msg string) error   { return TError{RPC_PAYLOAD_TO_BYTES, _msg} }
func ERROR_RPC_RESPONSE_TO_STRING(_msg string) error { return TError{RPC_RESPONSE_TO_STRING, _msg} }
func ERROR_RPC_RESPONSE_RESULT_TO_JSON(_msg string) error {
	return TError{RPC_RESPONSE_RESULT_TO_JSON, _msg}
}
func ERROR_RPC_RESPONSE_ERROR_TO_JSON(_msg string) error {
	return TError{RPC_RESPONSE_ERROR_TO_JSON, _msg}
}
func ERROR_RPC_REQUEST(_msg string) error { return TError{RPC_REQUEST, _msg} }

// api
func ERROR_API_UNSUPPORTED_ROUTER_METHOD() error {
	return TError{API_UNSUPPORTED_ROUTER_METHOD, MSG_API_UNSUPPORTED_ROUTER_METHOD}
}
