package apigee

import (
	"path"
)

// CacheService is an interface for interfacing with the Apigee Edge Admin API
// dealing with apiproducts.
type CacheService interface {
	Get(string, string) (*Cache, *Response, error)
	Create(Cache, string) (*Cache, *Response, error)
	Delete(string, string) (*Response, error)
	Update(Cache, string) (*Cache, *Response, error)
}

type CacheServiceOp struct {
	client *EdgeClient
}

var apiEndpoint = "caches"
var _ CacheService = &CacheServiceOp{}

type Cache struct {
	Name                              string `json:"name,omitempty"`
	Description                       string `json:"description,omitempty"`
	OverflowToDisk                    bool   `json:"overflowToDisk,omitempty"`
	SkipCacheIfElementSizeInKBExceeds string `json:"skipCacheIfElementSizeInKBExceeds,omitempty"`
}

func (s *CacheServiceOp) Get(name, env string) (*Cache, *Response, error) {

	path := path.Join("environments", env, apiEndpoint, name)

	req, e := s.client.NewRequest("GET", path, nil, "")
	if e != nil {
		return nil, nil, e
	}
	returnedCache := Cache{}
	resp, e := s.client.Do(req, &returnedCache)
	if e != nil {
		return nil, resp, e
	}
	return &returnedCache, resp, e

}

func (s *CacheServiceOp) Create(cache Cache, env string) (*Cache, *Response, error) {
	return postOrPutCache(cache, env, "POST", s)
}

func (s *CacheServiceOp) Update(cache Cache, env string) (*Cache, *Response, error) {
	return postOrPutCache(cache, env, "PUT", s)
}

func (s *CacheServiceOp) Delete(name, env string) (*Response, error) {
	path := path.Join("environments", env, apiEndpoint, name)

	req, e := s.client.NewRequest("DELETE", path, nil, "")
	if e != nil {
		return nil, e
	}

	resp, e := s.client.Do(req, nil)
	if e != nil {
		return resp, e
	}

	return resp, e

}

type CacheOptions struct {
	Name string `url:"name"`
}

func postOrPutCache(cache Cache, env, opType string, s *CacheServiceOp) (*Cache, *Response, error) {
	var err error
	uripath := ""

	if opType == "PUT" {
		uripath = path.Join("environments", env, apiEndpoint, cache.Name)
	} else {
		opt := CacheOptions{Name: cache.Name}
		uripath, err = addOptions(path.Join("environments", env, apiEndpoint), opt)
		if err != nil {
			return nil, nil, err
		}
	}

	req, e := s.client.NewRequest(opType, uripath, cache, "")
	if e != nil {
		return nil, nil, e
	}

	returnedCache := Cache{}

	resp, e := s.client.Do(req, &returnedCache)
	if e != nil {
		return nil, resp, e
	}

	return &returnedCache, resp, e
}
