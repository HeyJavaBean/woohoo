package stage

import (
	"sync"
)

//定义了函数节点链表的情况
type Stage struct {
	Input     *chan interface{}
	Output    *chan interface{}
	Func      StageFunc
	Next      *Stage
	Stateless bool
}

type StageFunc interface {
	Fire(in interface{}, output *chan interface{}, wg *sync.WaitGroup)
}

func (function *Stage) FireValve() {
	go func() {
		if function.Stateless {
			wg := sync.WaitGroup{}
			for {
				if data, ok := <-*(function.Input); ok {
					wg.Add(1)
					function.Func.Fire(data, function.Output, &wg)
					continue
				}
				wg.Wait()
				break
			}
		} else {
			pool := []interface{}{}
			for {
				if data, ok := <-*(function.Input); ok {
					pool = append(pool, data)
					continue
				}
				function.Func.Fire(pool, function.Output, &sync.WaitGroup{})
				break
			}
		}
		close(*function.Output)

	}()
}
