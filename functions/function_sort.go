package functions

import "sync"

//todo 想用布隆过滤器来做

//感觉这个有点问题

type ValveSort struct{
	mapper map[interface{}]bool
}


func (valve *ValveSort) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){

	if !valve.mapper[in]{
		valve.mapper[in]=true
		*output<-in
	}

	//一个不优雅的写法问题
	wg.Done()

}

func NewSort() *ValveSort {
		f:= new(ValveSort)
		f.mapper =  map[interface{}]bool{}
		return f
}