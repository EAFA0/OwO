package http

import "net/http"

type option struct {
	header http.Header
}

type Option func(*option)

// WithHeader 添加指定 header 值
func WithHeader(key, value string) Option {
	return func(o *option) {
		o.header[key] = append(o.header[key], value)
	}
}
