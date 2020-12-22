package stage

import (
	"sync"
)




type ValveThenGroup struct{
	identifyFunc IdentifyFunc
}

//k是string 但是 v 是结果数组！
type Entry struct{
	K string
	Vs []interface{}
}

//输出类型是一个自定义的entry类型
func (valve *ValveThenGroup) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){

	ins := in.([]interface{})

	res := map[string][]interface{}{}
	for _, data := range ins {
		k := valve.identifyFunc(data)
		res[k] = append(res[k], data)
	}


	//按照顺序输出
	for k, v := range res {
		*output<-Entry{k,v}
	}

}

func NewThenGroup(identifyFunc IdentifyFunc) *ValveThenGroup{
	f:= new(ValveThenGroup)
	f.identifyFunc = identifyFunc
	return f
}