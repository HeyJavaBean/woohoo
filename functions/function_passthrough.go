package functions

import "sync"

type ValvePassthrough struct{

}

func (valve *ValvePassthrough) doFire(in interface{},output *chan interface{},wg *sync.WaitGroup){
	*output<-in
	wg.Done()
}





func (valve *ValvePassthrough) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){
	go valve.doFire(in,output,wg)
}

func NewPassthrough() *ValvePassthrough{
	return new(ValvePassthrough)
}