package rest

import (
	"time"

	"github.com/chenhaonan-eth/go-ftx/auth"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

const ENDPOINT = "https://ftx.com/api"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Client struct {
	Auth *auth.Config

	HTTPC       *fasthttp.Client
	HTTPTimeout time.Duration
}

func New(auth *auth.Config, proxy string) *Client {
	hc := &fasthttp.Client{
		Dial: fasthttpproxy.FasthttpHTTPDialer(proxy),
	}
	if proxy == "" {
		hc = new(fasthttp.Client)
	}
	return &Client{
		Auth:        auth,
		HTTPC:       hc,
		HTTPTimeout: 5 * time.Second,
	}
}
