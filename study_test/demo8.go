package main

import (
	"fmt"
	"sync"
)

/**
 * @author: xaohuihui
 * @Path: go_study/study_test/demo8.go
 * @Description:
 * @datetime: 2022/6/22 9:54:48
 * software: GoLand
**/

func main() {

	wg := &sync.WaitGroup{}
	total := 0
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
		wg.Add(1)
		go func(i int) {
			total += i
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("total: %d sum: %d", total, sum)
}
