package functions

type Function struct{

	//这个函数的本体，指定类型然后启动
	FuncType string

	//内容表，去这里找到真正的函数
	Funcs Detail

	//这个函数是否是无状态的（用不用等数据齐了）
	Stateless bool

	//下一个函数
	NextFunc *Function
}

func (function *Function) Fire(in interface{}) interface{}{
	typeName := function.FuncType
	if typeName=="map"{
		return function.Funcs.MapFunc(in)
	}
	if typeName=="filter"{
		if function.Funcs.FilterFunc(in){
			return in
		}
		return nil
	}
	if typeName=="flatMap"{
		arr := function.Funcs.FlatMapFunc(in)
		return arr
	}
	return nil
}



type Detail struct{
	MapFunc MapFunc
	FlatMapFunc FlatMapFunc
	FilterFunc FilterFunc

}

