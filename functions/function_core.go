package functions

import "sync"

type Function struct {
	Input *chan interface{}

	Output *chan interface{}

	Func ValveFunc
	//下一个函数
	NextFunc *Function
}

type ValveFunc interface {
	Fire(in interface{}, output *chan interface{},wg *sync.WaitGroup)
}

func (function *Function) FireValve() {
	//开一个携程触发了就不管了
	go func() {
		wg := sync.WaitGroup{}
		for {
			if data, ok := <-*(function.Input); ok {
				//并发方式：来几个数据开几个线程处理
				wg.Add(1)
				//执行完了里面自己done
				go function.Func.Fire(data, function.Output,&wg)
			} else {
				//等所有协程执行完了再关闭，这里只有一个人操作关闭所以没事
				wg.Wait()
				close(*function.Output)
				break
			}
		}
	}()
}


func AddValve(input, output *chan interface{}, fun ValveFunc) *Function {
	function := new(Function)
	function.Func = fun
	function.Input = input
	function.Output = output
	return function
}
