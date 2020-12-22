package stage

import "sync"

//todo 想用布隆过滤器来做

//感觉这个有点问题

type ValveDistinct struct{
	mapper map[interface{}]bool
}


func (valve *ValveDistinct) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){

	if !valve.mapper[in]{
		valve.mapper[in]=true
		*output<-in
	}

	//一个不优雅的写法问题
	wg.Done()

}

func NewDistinct() *ValveDistinct {
		f:= new(ValveDistinct)
		f.mapper =  map[interface{}]bool{}
		return f
}