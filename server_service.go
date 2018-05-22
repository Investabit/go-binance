package binance

import (
	"context"
	"net/http"
)

// PingService ping server
type PingService struct {
	c *Client
}

// Do send request
func (s *PingService) Do(ctx context.Context, opts ...RequestOption) (raw *http.Response, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/ping",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	return data.Response, err
}

// ServerTimeService get server time
type ServerTimeService struct {
	c *Client
}

// Do send request
func (s *ServerTimeService) Do(ctx context.Context, opts ...RequestOption) (serverTime int64, raw *http.Response, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/time",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return 0, data.Response, err
	}
	j, err := newJSON(data.Data)
	if err != nil {
		return 0, data.Response, err
	}
	serverTime = j.Get("serverTime").MustInt64()
	return serverTime, data.Response, nil
}
