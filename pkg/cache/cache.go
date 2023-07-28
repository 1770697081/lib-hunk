package cache

import (
	"runtime/debug"
	"sync"

	"github.com/coocood/freecache"
)

var (
	CacheFd          *CacheStu
	muxCachePushTime sync.Mutex
)

func newFreecache(num int) *CacheStu {
	if num <= 0 {
		num = 10
	}
	pp := &CacheStu{}
	pp.Fd = freecache.NewCache(num * 1024 * 1024)
	debug.SetGCPercent(20)
	return pp
}

// 获取全局变量;
func NewFreeCacheGlobal(num int) *CacheStu {
	if CacheFd == nil {
		muxCachePushTime.Lock()
		if CacheFd == nil {
			CacheFd = newFreecache(num)
		}
		muxCachePushTime.Unlock()
	}
	return CacheFd
}

type CacheStu struct {
	Fd *freecache.Cache
}

func (m *CacheStu) SetBytes(key string, value []byte, timeout int64) (err error) {
	err = m.Fd.Set([]byte(key), []byte(value), int(timeout))
	return
}
func (m *CacheStu) SetString(key string, value string, timeout int64) (err error) {
	err = m.Fd.Set([]byte(key), []byte(value), int(timeout))
	return
}
func (m *CacheStu) GetString(key string) (byteData string, err error) {
	value, err := m.Fd.Get([]byte(key))
	byteData = string(value)
	return
}
func (m *CacheStu) GetBytes(key string) (value []byte, err error) {
	value, err = m.Fd.Get([]byte(key))
	return
}
