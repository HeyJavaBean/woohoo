package stage

import "sync"


type MeterCounter *int

//其实就是包装了一个int指针，在执行完流之后使用*a获取值就可以了
func NewCounter() MeterCounter{
	a := 0
	return &a
}


type ValveMeter struct{
	count *int
}


func (valve *ValveMeter) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){

	*output<-in
	*valve.count++
	wg.Done()

}

func NewMeter(count MeterCounter) *ValveMeter{
	f:= new(ValveMeter)
	f.count = count
	return f
}