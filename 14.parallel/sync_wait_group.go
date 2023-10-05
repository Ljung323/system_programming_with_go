package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // ジョブ数をあらかじめ登録
	wg.Add(2)
	go func() {
		// 非同期で仕事をする(1)
		fmt.Println(" 仕事 1")
		// Doneで完了を通知
		wg.Done()
	}()
	go func() {
		// 非同期で仕事をする(2)
		fmt.Println(" 仕事 2")
		// Doneで完了を通知
		wg.Done()
	}()
	// すべての処理が終わるのを待つ
	wg.Wait()
	fmt.Println(" 終了 ")
}
