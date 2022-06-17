package main

import (
	"sync"
	"sync/atomic"
)

/**
 * @author: xaohuihui
 * @Path: go_study/study_test/once_instance.go
 * @Description:
 * @datetime: 2022/6/2 18:38:58
 * software: GoLand
**/


var initialized uint32
type singleton struct {}
var instance *singleton
var mu *sync.Mutex

func GetInstance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}


type singletonOnce struct {}

var instanceOnce *singletonOnce
var once *sync.Once

func GetInstanceOnce() *singletonOnce {
	once.Do(func() {
		instanceOnce = &singletonOnce{}
	})
	return instanceOnce
}