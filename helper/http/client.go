package httpclient

import (
	"net/http"
	"time"

	httpcache "github.com/sniperkit/github-release-checker/helper/http/cache"
	httpstats "github.com/sniperkit/github-release-checker/helper/http/stats"

	"github.com/k0kubun/pp"
)

var (
	DefaultRequestDelayWorker time.Duration = (100 * time.Millisecond)
	DefaultTimeoutDial        time.Duration = (5 * time.Second)
	DefaultTimeoutRequest     time.Duration = (5 * time.Second)
	DefaultTimeoutResponse    time.Duration = (5 * time.Second)
)

var (
	defaultConfig    *Config
	defaultClient    = &http.Client{}
	defaultTransport = &http.Transport{}
)

var Loaded bool

func NewClient(cfg *Config) *http.Client {
	defaultConfig.Debug = cfg.Debug
	if cfg.Cache {
		defaultClient.Transport = httpcache.NewCacheWithTransport(defaultTransport)
		defaultConfig.Cache = true
	}
	if cfg.Stats {
		defaultClient.Transport = httpstats.NewTransport(defaultClient.Transport)
		defaultConfig.Stats = true
	}
	defaultConfig.healthy = true
	defaultConfig.loaded = true
	if cfg.Debug {
		pp.Println("custom http client: ", defaultConfig.Info())
	}

	return defaultClient
}

func NewTransport(cfg *Config) http.RoundTripper {
	client := &http.Client{}
	transport := &http.Transport{}
	if cfg.Cache {
		client.Transport = httpcache.NewCacheWithTransport(transport)
	}
	if cfg.Stats {
		client.Transport = httpstats.NewTransport(client.Transport)
	}
	return client.Transport
}

type Config struct {
	healthy bool
	loaded  bool
	Debug   bool
	Cache   bool
	Stats   bool
	Proxy   bool // not implemented yet
	Worker  Worker
	Timeout Timeout
}

func (tc *Config) Info() map[string]bool {
	info := make(map[string]bool, 3)
	info["pkg.loaded"] = Loaded
	info["client.loaded"] = tc.loaded
	info["client.healthy"] = tc.healthy
	info["client.debug"] = tc.Debug
	info["client.cache"] = tc.Cache
	info["client.stats"] = tc.Stats
	info["client.proxy"] = tc.Proxy
	return info
}

func (c *Config) Reload(cfg *Config) *http.Client {
	c = cfg
	return NewClient(cfg)
}

type Worker struct {
	Delay time.Duration
}

type Timeout struct {
	Dial     time.Duration
	Request  time.Duration
	Response time.Duration
}

func init() {
	Loaded = true
	defaultConfig = &Config{}
}
