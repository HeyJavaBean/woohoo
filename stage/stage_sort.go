package stage

import "sync"

//todo 想用布隆过滤器来做


type ValveSort struct{
	comparator Comparator
	desc bool
}


func (valve *ValveSort) Fire(in interface{},output *chan interface{},wg *sync.WaitGroup){

	//挨个排序，直到完成了才放行
	ins := in.([]interface{})

	//利用函数进行排序,经典泡泡排序
	for i:=0;i<len(ins)-1;i++{

		for j:=i+1;j<len(ins);j++{

			flag := valve.comparator(ins[i], ins[j])

			if valve.desc{
				flag = !flag
			}

			if flag {
				temp := ins[i]
				ins[i] = ins[j]
				ins[j] = temp
			}
		}

	}

	//按照顺序输出
	for _, i := range ins {
		*output<-i
	}


}

func NewSort(comparator Comparator, desc bool) *ValveSort {
		f:= new(ValveSort)
		f.comparator = comparator
		return f
}