package functions

import "sync"


type ValveLimit struct{
	limitNum int
}

//思路：来了的放到channel里，channel出口是单线程搬运，就安全了
func (valve *ValveLimit) pass(output *chan interface{}){

}

func (valve *ValveLimit) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){

	//todo 这里看看怎么解决
	//保证线程安全的同时进行计数
	//out := valve.MapFunc(in)
	//if out!=nil{
	//	*output<-out
	//}
	//wg.Done()

}

func NewLimit(limitNum int) *ValveLimit {
		f:= new(ValveLimit)
		f.limitNum = limitNum
		return f
}