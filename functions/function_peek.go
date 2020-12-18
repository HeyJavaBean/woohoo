package functions

import "sync"

type PeekFunc func(interface{})

type ValvePeek struct{
	PeekFunc PeekFunc
}


func (valve *ValvePeek) doFire(in interface{},output *chan interface{},wg *sync.WaitGroup){
	//为了让后续更快的执行，所以先丢给后面，自己再慢慢打印
	*output<-in
	valve.PeekFunc(in)
	wg.Done()
}






func (valve *ValvePeek) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){
	go valve.doFire(in,output,wg)
}

func NewPeek(peekFunc PeekFunc) *ValvePeek{
		f:= new(ValvePeek)
		f.PeekFunc = peekFunc
		return f
}