package main

import (
	"fmt"
	"sync"
)

var lock=&sync.Mutex{}

type single struct {

}
var singleInstance *single

func getInstance() *single {
	if singleInstance !=nil {
		fmt.Println("single instance already exist")
		return singleInstance
	}
	lock.Lock()

	defer lock.Unlock()

	if singleInstance  ==nil {
		fmt.Println("creating single instance now")
		singleInstance=&single{}
	}
	fmt.Println("single instance already created")
	return singleInstance

}
