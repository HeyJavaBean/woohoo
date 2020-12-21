package stage

import "sync"

type ValveSafePeek struct{
	PeekFunc PeekFunc
}


//特点，单线程执行，所以是安全的
func (valve *ValveSafePeek) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){
	*output<-in
	valve.PeekFunc(in)
	wg.Done()
}

func NewSafePeek(peekFunc PeekFunc) *ValveSafePeek{
		f:= new(ValveSafePeek)
		f.PeekFunc = peekFunc
		return f
}