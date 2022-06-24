package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

/* 控制协程数量计算slice 字母出现的个数 */

func sumCount(str string, gChan chan struct{}, wg *sync.WaitGroup, resChan chan map[rune]int) {
	<-gChan
	var res = make(map[rune]int)
	defer func() {
		wg.Done()
	}()
	for _, ch := range str {
		res[ch] += 1
	}
	resChan <- res
}

func main() {
	_set := make([]string, 10, 10)
	_set = append(_set, []string{"sce", "sa", "g你re", "sdfa", "你grereee", "sw好eq"}...)
	var gChan = make(chan struct{}, 2)
	var resChan = make(chan map[rune]int, len(_set))
	defer func() {
		defer close(gChan)
		defer close(resChan)
	}()

	wg := &sync.WaitGroup{}
	for i, _ := range _set {
		wg.Add(1)
		gChan <- struct{}{}
		go sumCount(_set[i], gChan, wg, resChan)
	}

	res := map[rune]int{}

	for i := 0; i < len(_set); i++ {
		tmp := <-resChan
		for k, v := range tmp {
			res[k] += v
		}

	}

	for k, v := range res {
		fmt.Printf("k: %s v: %d \n", string(k), v)
	}
	wg.Wait()

	go func() {
		for {
			log.Printf("len: %d", Add("go-programming-tour-book"))
			time.Sleep(time.Microsecond * 10)
		}
	}()
	_ = http.ListenAndServe("0.0.0.0:6060", nil)

}

var datas []string
func Add(str string) int {
	data := []byte(str)
	datas  = append(datas, string(data))
	return len(datas)
}