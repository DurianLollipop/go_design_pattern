// 创建型模式：单例模式

package singleton

import "sync"

// 单例模式结构体
type Singleton struct{}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
