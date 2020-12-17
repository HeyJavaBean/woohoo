package functions

import "sync"

type ValvePassthrough struct{

}

func (valve *ValvePassthrough) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){


	*output<-in
	wg.Done()


}

func NewPassthrough() *ValvePassthrough{
	return new(ValvePassthrough)
}