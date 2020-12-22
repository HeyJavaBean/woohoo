package stage

import "sync"

type FlatMapFunc func(interface{}) []interface{}



type ValveFlatMap struct{
	FlatMapFunc FlatMapFunc
}

func (valve *ValveFlatMap) doFire(in interface{},output *chan interface{},wg *sync.WaitGroup){

	arr := valve.FlatMapFunc(in)
	if arr!=nil{
		for _,a := range arr {
			*output<-a
		}
	}
	wg.Done()

}



func (valve *ValveFlatMap) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){
	go valve.doFire(in,output,wg)
}

func NewFlatMap(fmap FlatMapFunc) *ValveFlatMap{
	f:= new(ValveFlatMap)
	f.FlatMapFunc = fmap
	return f
}