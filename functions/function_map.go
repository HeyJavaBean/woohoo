package functions

import "sync"

type MapFunc func(interface{}) interface{}

type ValveMap struct{
	MapFunc MapFunc
}

func (valve *ValveMap) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){
	out := valve.MapFunc(in)
	if out!=nil{
		*output<-out
	}
	wg.Done()
}

func NewMap(mapFunc MapFunc) *ValveMap{
		f:= new(ValveMap)
		f.MapFunc = mapFunc
		return f
}