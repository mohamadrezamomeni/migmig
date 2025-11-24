package migmig

import (
	"context"
	"fmt"
	"net/http"
)

type Config struct {
	Headers map[string]string `json:"headers,omitempty"`
}

func CreateConfig() *Config {
	return &Config{
		Headers: make(map[string]string),
	}
}

type Demo struct {
	next    http.Handler
	headers map[string]string
	name    string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Headers) == 0 {
		return nil, fmt.Errorf("headers cannot be empty")
	}

	return &Demo{
		headers: config.Headers,
		next:    next,
		name:    name,
	}, nil
}

func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hellllllo"))
}
