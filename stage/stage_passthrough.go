package stage

import "sync"

type ValvePassthrough struct{

}





//保证直接输出的熟悉，不能并发
func (valve *ValvePassthrough) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){
	*output<-in
	wg.Done()
}

func NewPassthrough() *ValvePassthrough{
	return new(ValvePassthrough)
}