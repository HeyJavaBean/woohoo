package stage

import "sync"



type ValveMap struct{
	MapFunc MapFunc
}

func (valve *ValveMap) doFire(in interface{},output *chan interface{},wg *sync.WaitGroup){

	out := valve.MapFunc(in)
	if out!=nil{
		*output<-out
	}
	wg.Done()

}



func (valve *ValveMap) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){
	go valve.doFire(in,output,wg)
}

func NewMap(mapFunc MapFunc) *ValveMap{
		f:= new(ValveMap)
		f.MapFunc = mapFunc
		return f
}