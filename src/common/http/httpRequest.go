// Created by Seunggwan, Back on 2024.04.24
// Copyright (C) 2022-2024 Seunggwan, Back - All Rights Reserved
package http

import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"gorani/common/errors"
	"gorani/common/utils"
	jMapper "gorani/common/utils/json"
)

// NewRequest는 지정된 메소드에 기반한 HTTP 요청을 생성합니다.
func NewRequest(_method, _url string, _body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(_method, _url, _body)
	if err != nil {
		return nil, errors.ERROR_HTTP_NEW_REQUEST(err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Close = true
	return req, nil
}

func NewGetRequest(_url string, _body io.Reader) (*http.Request, error) {
	return NewRequest(GET.String(), _url, _body)
}

func NewPostRequest(_url string, _body io.Reader) (*http.Request, error) {
	return NewRequest(POST.String(), _url, _body)
}

// NewHttpClient는 지정된 타임아웃과 안전하지 않은 TLS 설정으로 새 HTTP 클라이언트를 생성합니다.
func NewHttpClient(_timeout int) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 주의: 이것은 안전하지 않습니다; 주의해서 사용하세요!
		},
		Timeout: time.Millisecond * time.Duration(_timeout),
	}
}

// sendRequest는 HTTP 요청을 보내고 응답 본문을 바이트 슬라이스로 반환합니다.
func sendRequest(_req *http.Request, _client *http.Client) ([]byte, error) {
	resp, err := _client.Do(_req)
	if err != nil {
		return nil, errors.ERROR_HTTP_SEND_REQUEST(err.Error())
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// PostRequestToBytes는 POST 요청을 보내고 응답을 바이트로 반환합니다.
func PostRequestToBytes(_url, _body string, _timeout int) ([]byte, error) {
	req, err := NewPostRequest(_url, bytes.NewBufferString(_body))
	if err != nil {
		return nil, err
	}

	client := NewHttpClient(_timeout)
	return sendRequest(req, client)
}

// PostRequestFromString은 POST 요청을 보내고 응답을 JSON 맵으로 구문 분석합니다.
func PostRequestFromString(_url, _body string, _timeout int) (*jMapper.TJsonMap, error) {
	bytes, err := PostRequestToBytes(_url, _body, _timeout)
	if err != nil {
		return nil, err
	}

	jmap, err := jMapper.NewBytes(bytes)
	if err != nil {

		return nil, err
	}

	return jmap, nil
}

// PostRequestFromBytes는 바이트 슬라이스를 본문으로 사용하여 POST 요청을 보내고 응답을 JSON 맵으로 구문 분석합니다.
func PostRequestFromBytes(_url string, _byteBody []byte, _timeout int) (*jMapper.TJsonMap, error) {
	req, err := NewPostRequest(_url, bytes.NewBuffer(_byteBody))
	if err != nil {
		return nil, err
	}

	client := NewHttpClient(_timeout)
	bytes, err := sendRequest(req, client)
	if err != nil {
		return nil, err
	}

	return jMapper.NewBytes(bytes)
}

// GetRequest는 GET 요청을 보내고 응답을 bytes로 내보냅니다.
func GetRequestToBytes(_url string, _timeout int) ([]byte, error) {
	req, err := NewGetRequest(_url, nil)
	if err != nil {
		return nil, err
	}

	client := NewHttpClient(_timeout)
	bytes, err := sendRequest(req, client)
	if err != nil {
		return nil, err
	}

	return bytes, err
}

func GetRequestToString(_url string, _timeout int) (string, error) {
	bytes, err := GetRequestToBytes(_url, _timeout)
	if err != nil {
		return "", err
	}

	return utils.BytesToString(bytes), nil
}

func GetRequestToJson(_url string, _timeout int) (*jMapper.TJsonMap, error) {
	bytes, err := GetRequestToBytes(_url, _timeout)
	if err != nil {
		return nil, err
	}

	return jMapper.NewBytes(bytes)
}
