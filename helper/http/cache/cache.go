package http_cache

import (
	"net/http"
	"path/filepath"

	// c "github.com/sniperkit/cache"
	c "github.com/gregjones/httpcache"
	cb "github.com/sniperkit/cache/backend/local/badger"

	"github.com/sniperkit/github-release-checker/utils/fsutil"
)

var (
	loaded    bool
	HttpCache *c.Transport
)

func NewCacheWithTransport(ht http.RoundTripper) *c.Transport {
	cacheStoragePrefixPath := filepath.Join("shared", "cache", "badger")
	// cacheStoragePrefixPath := filepath.Join("/Users/lucmichalski/local/golang/src/github.com/roscopecoltran/sniperkit-agent/shared/data/cache/http", "cacher.badger")
	fsutil.EnsureDir(cacheStoragePrefixPath)
	hcache, err := cb.New(
		&cb.Config{
			ValueDir:    "api.github.com.v3.snappy",
			StoragePath: cacheStoragePrefixPath,
			SyncWrites:  true,
			Debug:       false,
			Compress:    true,
		})
	if err != nil {
		panic(err)
	}

	t := c.NewTransport(hcache)
	t.MarkCachedResponses = true
	// t.Debug = false
	t.Transport = ht
	return t
}

func NewDefaultCacheTransport() *c.Transport {
	cacheStoragePrefixPath := filepath.Join("shared", "cache", "badger")
	// cacheStoragePrefixPath := filepath.Join("/Users/lucmichalski/local/golang/src/github.com/roscopecoltran/sniperkit-agent/shared/data/cache/http", "cacher.badger")
	fsutil.EnsureDir(cacheStoragePrefixPath)
	hcache, err := cb.New(
		&cb.Config{
			ValueDir:    "api.github.com.v3.snappy",
			StoragePath: cacheStoragePrefixPath,
			SyncWrites:  true,
			Debug:       false,
			Compress:    true,
		})
	if err != nil {
		panic(err)
	}

	t := c.NewTransport(hcache)
	t.MarkCachedResponses = true
	// t.Debug = false
	return t
}

func NewCacheTransport(hrt http.RoundTripper) *c.Transport {
	cacheStoragePrefixPath := filepath.Join("shared", "cache", "badger")
	// cacheStoragePrefixPath := filepath.Join("/Users/lucmichalski/local/golang/src/github.com/roscopecoltran/sniperkit-agent/shared/data/cache/http", "cacher.badger")
	fsutil.EnsureDir(cacheStoragePrefixPath)
	hcache, err := cb.New(
		&cb.Config{
			ValueDir:    "api.github.com.v3.snappy",
			StoragePath: cacheStoragePrefixPath,
			SyncWrites:  true,
			Debug:       false,
			Compress:    true,
		})
	if err != nil {
		panic(err)
	}

	t := c.NewTransport(hcache)
	t.MarkCachedResponses = true
	t.Transport = hrt
	// t.Debug = false
	return t
}

func IsLoaded() bool {
	return loaded
}

/*
func init() {
	HttpCache = NewDefaultCacheTransport()
}
*/
