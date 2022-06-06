package cache

import (
	"time"
)

type Cache struct {
	ma map[string]Data
}

type Data struct {
	val         string
	hasDeadline bool
	deadline    time.Time
}

func NewCache() *Cache {
	return &Cache{
		make(map[string]Data),
	}
}

func (cash *Cache) Put(key, value string) {
	data := Data{
		val:         value,
		hasDeadline: false,
	}
	cash.ma[key] = data
}

func (cash *Cache) PutTill(key, value string, deadline time.Time) {
	data := Data{
		val:         value,
		hasDeadline: true,
		deadline:    deadline,
	}
	cash.ma[key] = data
}

func (cash *Cache) Get(key string) (string, bool) {
	startTime := time.Now()

	data, ok := cash.ma[key]
	if !ok {
		return "", ok
	}

	if data.hasDeadline {
		notExp := startTime.Before(data.deadline)
		if notExp {
			return data.val, true
		} else {
			delete(cash.ma, key)
			return "", false
		}
	}
	return data.val, ok
}

func (cash *Cache) Keys() []string {
	var array []string
	for key := range cash.ma {
		array = append(array, key)

	}
	return array
}
