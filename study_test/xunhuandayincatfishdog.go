package main

import (
	"fmt"
	"sync"
)

// author: xaohuihui
// datetime: 2022/2/15 11:37:07
// software: GoLand

func Cat(dog, cat chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	<-dog
	fmt.Printf("cat ")
	cat <- struct{}{}
	defer func() {
		wg.Done()
	}()
}

func Fish(cat, fish chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	<-cat
	fmt.Printf("fish ")
	fish <- struct{}{}
	defer func() {
		wg.Done()
	}()

}

func Dog(fish, dog chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	<-fish
	fmt.Println("dog")
	dog <- struct{}{}
	defer func() {
		wg.Done()
	}()
}

func main() {
	var cat = make(chan struct{})
	var fish = make(chan struct{})
	var dog = make(chan struct{}, 1)
	var wg = &sync.WaitGroup{}

	defer func() {
		defer close(dog)
		defer close(cat)
		defer close(fish)
	}()

	dog <- struct{}{}
	for i := 0; i < 100; i++ {
		go Dog(fish, dog, wg)
		go Cat(dog, cat, wg)
		go Fish(cat, fish, wg)
	}

	wg.Wait()

}
