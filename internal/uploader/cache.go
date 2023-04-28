package uploader

import (
	"fmt"
	"strconv"
	"time"

	"github.com/duc-cnzj/mars/v4/internal/cache"
	"github.com/duc-cnzj/mars/v4/internal/contracts"
)

type cacheUploader struct {
	contracts.Uploader
	cacheFn func() contracts.CacheInterface
}

func NewCacheUploader(uploader contracts.Uploader, cache func() contracts.CacheInterface) contracts.Uploader {
	return &cacheUploader{Uploader: uploader, cacheFn: cache}
}

func toByteNum(i int64) []byte {
	return []byte(fmt.Sprintf("%d", i))
}

func byteNum(remember []byte) int64 {
	atoi, _ := strconv.Atoi(string(remember))
	return int64(atoi)
}

func (ca *cacheUploader) UnWrap() contracts.Uploader {
	return ca.Uploader
}

var DirSizeCacheSeconds = int((15 * time.Minute).Seconds())

func (ca *cacheUploader) DirSize() (int64, error) {
	remember, err := ca.cacheFn().Remember(cache.NewKey("dir-size"), DirSizeCacheSeconds, func() ([]byte, error) {
		size, err := ca.Uploader.DirSize()
		return toByteNum(size), err
	})
	return byteNum(remember), err
}
