package shell

import (
	"bufio"
	"io"
	"net/http"
)

type IOStream struct {
	writer []io.Writer
	reader []io.Reader
}

type SystemIO struct {
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
