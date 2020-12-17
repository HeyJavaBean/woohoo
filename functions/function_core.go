package functions

import (
	"fmt"
)

type Function struct {
	Input *chan interface{}

	Output *chan interface{}

	Func ValveFunc
	//下一个函数
	NextFunc *Function
}

type ValveFunc interface {
	Fire(in interface{}, output *chan interface{})
}

func (function *Function) FireValve() {
	fmt.Println("fire!")
	go func() {
		for {
			if data, ok := <-*(function.Input); ok {
				function.Func.Fire(data, function.Output)
			} else {
				//关闭输出管道?
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
