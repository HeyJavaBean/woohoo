package functions


type FlatMapFunc func(interface{}) []interface{}



type ValveFlatMap struct{
	FlatMapFunc FlatMapFunc
}

func (valve *ValveFlatMap) Fire(in interface{},output *chan interface{}){

	arr := valve.FlatMapFunc(in)
	if arr!=nil{
		for _,a := range arr {
			*output<-a
		}
	}

}

func NewFlatMap(fmap FlatMapFunc) *ValveFlatMap{
	f:= new(ValveFlatMap)
	f.FlatMapFunc = fmap
	return f
}