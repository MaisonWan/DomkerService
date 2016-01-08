package watcher

import (
	"fmt"
	"time"
)

// 监控类
type Watcher struct {
}

func (w *Watcher) Start() {
	fmt.Println("start wathcer.")
}

func (w *Watcher) Stop() {
	fmt.Println("stop watcher.")
}

func Say() {
	fmt.Println("now is ", time.Now())
}
