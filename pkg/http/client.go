package http

import (
	"context"
	"net/http"
)

func Get(ctx context.Context, url string, body interface{}, opts ...Option) (*http.Response, error) {
	return Request(ctx, "GET", url, body, opts...)
}
func Post(ctx context.Context, url string, body interface{}, opts ...Option) (*http.Response, error) {
	return Request(ctx, "POST", url, body, opts...)
}
func Put(ctx context.Context, url string, body interface{}, opts ...Option) (*http.Response, error) {
	return Request(ctx, "PUT", url, body, opts...)
}
func Delete(ctx context.Context, url string, body interface{}, opts ...Option) (*http.Response, error) {
	return Request(ctx, "DELETE", url, body, opts...)
}

func Request(ctx context.Context, method, url string, body interface{}, opts ...Option) (*http.Response, error) {
	return nil, nil
}
