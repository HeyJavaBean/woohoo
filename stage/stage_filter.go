package stage

import "sync"






type ValveFilter struct{
	FilterFunc FilterFunc
}




func (valve *ValveFilter) doFire(in interface{},output *chan interface{},wg *sync.WaitGroup){

	out := valve.FilterFunc(in)
	if out{
		*output<-in
	}
	wg.Done()
}

func (valve *ValveFilter) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){
	go valve.doFire(in,output,wg)
}

func NewFilter(filterFunc FilterFunc) *ValveFilter{
	f:= new(ValveFilter)
	f.FilterFunc = filterFunc
	return f
}