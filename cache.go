package cache

import "errors"

type Cache interface {
	New()
}

type Data struct {
	cache map[string]interface{}
}

func New() Data {
	var newCacheObject = make(map[string]interface{})
	return Data{newCacheObject}
}

func (cache Data) Set(key string, value interface{}) error {
	if len(key) == 0 {
		return errors.New("the key length is zero")
	} else if value == nil {
		return errors.New("the value is nil")
	}

	cache.cache[key] = value
	return nil
}

func (cache Data) Get(key string) (interface{}, error) {
	err := checkZero(key)

	return cache.cache[key], err
}

func (cache Data) Delete(key string) error {
	checkZero(key)

	delete(cache.cache, key)
	return nil
}

func checkZero(key string) error{
	if len(key) == 0 {
		return errors.New("the key length is zero")
	}
	return nil
}
