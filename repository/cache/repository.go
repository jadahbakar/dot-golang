package cache

import (
	"errors"
	"net/url"

	"github.com/go-redis/redis/v9"
	"github.com/jadahbakar/dot-golang/domain/bayar"
	"github.com/jadahbakar/dot-golang/domain/siswa"
)

type redisRepository struct {
	cache *redis.Client
}

type InitFunc func(url *url.URL) (*redis.Client, error)

var cacheImpl = make(map[string]InitFunc)

func newClient(urlStr string) (*redis.Client, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	f, ok := cacheImpl[u.Scheme]
	if !ok {
		return nil, errors.New("unsupported scheme")
	}

	return f(u)
}

func NewRedis(urlStr string) (siswa.Repository, bayar.Repository, error) {
	drv, err := newClient(urlStr)
	if err != nil {
		return nil, nil, err
	}
	siswaRepo := &redisRepository{cache: drv}
	bayarRepo := &redisRepository{cache: drv}
	return siswaRepo, bayarRepo, nil
}

func (r *redisRepository) GetAllSiswa(nis string) (siswa.Siswa, error) {  {

	r, err:= r.cache.Get(ctx, nis)
	
	return nil,nil
}