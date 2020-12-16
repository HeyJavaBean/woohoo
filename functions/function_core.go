package functions

type Function struct{
	Func ValveFunc
	//下一个函数
	NextFunc *Function
}

type ValveFunc interface {
	Fire(in interface{}) []interface{}
}

func AddValve(fun ValveFunc) *Function{
	function := new(Function)
	function.Func = fun
	return function
}