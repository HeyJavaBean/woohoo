package functions

import "sync"

//todo 想用布隆过滤器来做

//感觉这个有点问题
type Comparator func(a,b interface{}) bool


type ValveSort struct{
	comparator Comparator
	sortList []interface{}
}


func (valve *ValveSort) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){

	//挨个排序，直到完成了才放行

	if !valve.mapper[in]{
		valve.mapper[in]=true
		*output<-in
	}

	//一个不优雅的写法问题
	wg.Done()

}

func NewSort(comparator Comparator) *ValveSort {
		f:= new(ValveSort)
		f.comparator = comparator
		return f
}