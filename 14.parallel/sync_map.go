package main

import (
	"fmt"
	"sync"
)

func main() {
	// 初期化
	smap := &sync.Map{}

	// なんでも入れられる
	smap.Store("hello", "world")
	smap.Store(1, 2)

	// 削除
	smap.Delete("test")

	// 取り出し方法
	value, ok := smap.Load("hello")
	fmt.Printf("key=%v value=%v exists?=%v\n", "hello", value, ok)

	// 全要素を取り出す
	printAll(smap)

	// これはすでに保存されているので無視
	smap.LoadOrStore(1, 3)
	printAll(smap)

	// これは保存されていないので挿入
	smap.LoadOrStore(2, 4)
	printAll(smap)
}

func printAll(smap *sync.Map) {
	fmt.Println("====")
	smap.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})
	fmt.Println("====")
}
