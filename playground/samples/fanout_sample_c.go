package samples

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func repeatFunc(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn(): // 関数の実行結果をvalueStreamに格納
			}
		}
	}()
	return valueStream
}
func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ { // これは並行処理    num回並行してvalueStreamからの値を受け取る
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}
func toInt(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for v := range valueStream { // interface型の値をintに型アサーションする
			select {
			case <-done:
				return
			case intStream <- v.(int):
			}
		}
	}()
	return intStream
}
func primeFinder(done <-chan interface{}, intStream <-chan int) <-chan interface{} { // 素数を探す
	primeStream := make(chan interface{})
	go func() {
		defer close(primeStream)
	L:
		for i := range intStream {
			for div := 2; div < i; div++ {
				if i%div == 0 {
					continue L
				}
			}
			select {
			case <-done:
				return
			case primeStream <- i:
				fmt.Println("result!")
			}
		}
	}()
	return primeStream
}
func random() interface{} { // 素数を探すための前準備(ランダムに整数生成)
	return rand.Intn(500000000)
}

func fanInxxx(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})
	multiplex := func(c <-chan interface{}) { // 並行して各チャネルがmultiplexedStreamチャネルに値を入れる
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}
	for _, c := range channels { // 非同期に複数(CPU数に応じて)のチャネルを起動してそれぞれにmultiplexを実行させる
		wg.Add(1)
		go multiplex(c)
	}
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()
	return multiplexedStream
}
func SampleFanoutBBB() {
	// 処理数十秒かかる、ランダム値から素数を10個取得する処理からfan-out/fan-inを使って処理時間を改善する
	done := make(chan interface{})
	defer close(done)
	randIntStream := toInt(done, repeatFunc(done, random)) // ランダム値をrepeatで何度も生成してint型にしてrandIntStreamチャネルにわたす
	start := time.Now()

	// fan in out使った例
	numFinders := runtime.NumCPU()
	fmt.Printf("prime finders: %v\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	output := fanInxxx(done, finders...)
	for prime := range take(done, output, 5) {
		fmt.Println(prime)
	}

	// fan in out使わない例(時間かかる)
	// for prime := range take(done, primeFinder(done, randIntStream), 10) { // 10個素数を見つける
	// 	fmt.Println(prime)
	// }

	fmt.Println(time.Since(start))
}
