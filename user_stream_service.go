package binance

import (
	"context"
	"net/http"
)

// StartUserStreamService create listen key for user stream service
type StartUserStreamService struct {
	c *Client
}

// Do send request
func (s *StartUserStreamService) Do(ctx context.Context, opts ...RequestOption) (listenKey string, raw *http.Response, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/api/v1/userDataStream",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", data.Response, err
	}
	j, err := newJSON(data.Data)
	if err != nil {
		return "", data.Response, err
	}
	listenKey = j.Get("listenKey").MustString()
	return listenKey, data.Response, nil
}

// KeepaliveUserStreamService update listen key
type KeepaliveUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *KeepaliveUserStreamService) ListenKey(listenKey string) *KeepaliveUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveUserStreamService) Do(ctx context.Context, opts ...RequestOption) (raw *http.Response, err error) {
	r := &request{
		method:   "PUT",
		endpoint: "/api/v1/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	data, err := s.c.callAPI(ctx, r, opts...)
	return data.Response, err
}

// CloseUserStreamService delete listen key
type CloseUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *CloseUserStreamService) ListenKey(listenKey string) *CloseUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *CloseUserStreamService) Do(ctx context.Context, opts ...RequestOption) (raw *http.Response, err error) {
	r := &request{
		method:   "DELETE",
		endpoint: "/api/v1/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	data, err := s.c.callAPI(ctx, r, opts...)
	return data.Response, err
}
