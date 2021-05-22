package io

import (
	"bufio"
	"io"
	"net/http"
)

type Stream struct {
	Writer []io.Writer
	Reader []io.Reader
}

type Buffer struct {
	writer bufio.Writer
	reader bufio.Reader
}

type HttpServer struct {
	writer http.ResponseWriter
	reader http.HandlerFunc
	server http.Server
}
type HttpClient struct {
	writer io.Writer
	reader io.Reader
	client http.Client
}
