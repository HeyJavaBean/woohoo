package stage

import (
	"sync"
)




type ValveThenGroupCount struct{
	identifyFunc IdentifyFunc
}

//k是string 但是 v 是结果数组！
type EntryCount struct{
	K string
	V int
}

//输出类型是一个自定义的entry类型
func (valve *ValveThenGroupCount) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){

	ins := in.([]interface{})

	res := map[string]int{}
	for _, data := range ins {
		k := valve.identifyFunc(data)
		res[k]++
	}


	//按照顺序输出
	for k, v := range res {
		*output<-EntryCount{k,v}
	}

}

func NewThenGroupCount(identifyFunc IdentifyFunc) *ValveThenGroupCount{
	f:= new(ValveThenGroupCount)
	f.identifyFunc = identifyFunc
	return f
}