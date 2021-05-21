package shell

import (
	"bufio"
	"net/http"
	"os"
)

type StreamIO struct {
	buf  SystemIO
	file []*os.File
	rest Rest
}

type SystemIO struct {
	writer []*bufio.Writer
	reader []*bufio.Reader
}

type Rest struct {
	writer []*http.ResponseWriter
	reader []*http.HandlerFunc
	server []*http.Server
}
