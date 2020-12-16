package functions


type FlatMapFunc func(interface{}) []interface{}



type ValveFlatMap struct{
	FlatMapFunc FlatMapFunc
}

func (valve *ValveFlatMap) Fire(in interface{}) []interface{}{

	return valve.FlatMapFunc(in)

}

func NewFlatMap(fmap FlatMapFunc) *ValveFlatMap{
	f:= new(ValveFlatMap)
	f.FlatMapFunc = fmap
	return f
}