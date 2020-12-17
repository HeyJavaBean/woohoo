package functions

import "sync"

type FilterFunc func(interface{}) bool



type ValveFilter struct{
	FilterFunc FilterFunc
}

func (valve *ValveFilter) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){

	out := valve.FilterFunc(in)
	if out{
		*output<-in
	}
	wg.Done()
}

func NewFilter(filterFunc FilterFunc) *ValveFilter{
	f:= new(ValveFilter)
	f.FilterFunc = filterFunc
	return f
}