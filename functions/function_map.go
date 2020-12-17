package functions


type MapFunc func(interface{}) interface{}

type ValveMap struct{
	MapFunc MapFunc
}

func (valve *ValveMap) Fire(in interface{},output *chan interface{}){
	out := valve.MapFunc(in)
	if out!=nil{
		*output<-out
	}
}

func NewMap(mapFunc MapFunc) *ValveMap{
		f:= new(ValveMap)
		f.MapFunc = mapFunc
		return f
}