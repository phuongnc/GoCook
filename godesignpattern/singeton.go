package main

import "sync"

type Singleton struct{}

var (
	instanceSingleton1 *Singleton
	instanceSingleton2 *Singleton
	instanceSingleton3 *Singleton
	once               sync.Once
)

// 1.Init and get singleton1 by call init()
var singleton1 *Singleton

func init() {
	instanceSingleton1 = new(Singleton)
}

func GetInstance1() *Singleton {
	return singleton1
}

// 2. Get singleton2 by check nil
func GetInstance2() *Singleton {
	if instanceSingleton2 == nil {
		instanceSingleton2 = &Singleton{}
	}
	return instanceSingleton2
}

// 3. Get singleton3 by using sync.Once
func GetInstance3() *Singleton {
	once.Do(func() {
		instanceSingleton3 = new(Singleton)
	})
	return instanceSingleton3
}
