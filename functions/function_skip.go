package functions

import "sync"


type ValveSkip struct{
	skipNum int
}


func (valve *ValveSkip) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){

	if valve.skipNum<1{
		*output<-in
	}else{
		valve.skipNum--
	}

	//一个不优雅的写法问题
	wg.Done()

}

func NewSkip(skipNum int) *ValveLimit {
		f:= new(ValveLimit)
		f.limitNum = skipNum
		return f
}