package http_stats

import (
	"net/http"

	// http stats
	s "github.com/segmentio/stats"
	hs "github.com/segmentio/stats/httpstats"
	idb "github.com/segmentio/stats/influxdb"
)

var (
	loaded      bool
	StatsEngine s.Engine
)

var (
	idbCfg idb.ClientConfig
	idbCli *idb.Client
)

func NewTransport(rt http.RoundTripper) http.RoundTripper {
	return hs.NewTransport(rt)
}

func NewClientWithTransport(t *http.Transport) *http.Client {
	return &http.Client{
		Transport: hs.NewTransport(
			t,
		),
	}
}

func IsLoaded() bool {
	return loaded
}

/*
func init() {
	http.DefaultClient.Transport = hs.NewTransport(http.DefaultClient.Transport)
}
*/
