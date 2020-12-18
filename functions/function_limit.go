package functions

import "sync"


type ValveLimit struct{
	limitNum int
}


func (valve *ValveLimit) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){

	if valve.limitNum>0{
		*output<-in
		valve.limitNum--
	}

	//一个不优雅的写法问题
	wg.Done()

}

func NewLimit(limitNum int) *ValveLimit {
		f:= new(ValveLimit)
		f.limitNum = limitNum
		return f
}