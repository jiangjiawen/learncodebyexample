//优先从private空间拿，再加锁从shared空间拿，还没有再从其他的PoolLocal的shared空间拿，还没有就直接new一个返回。
//https://juejin.im/post/5d4087276fb9a06adb7fbe4a
// 来自蔡超谈软件
package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	pool.Put(1)
	pool.Put(2)
	pool.Put(3)
	pool.Put(4)
	for i := 0; i < 6; i++ {
		fmt.Println(pool.Get())
	}
}
