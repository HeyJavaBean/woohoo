package functions


type MapFunc func(interface{}) interface{}

type ValveMap struct{
	MapFunc MapFunc
}

func (valve *ValveMap) Fire(in interface{}) []interface{}{
	return []interface{}{valve.MapFunc(in)}
}

func NewMap(mapFunc MapFunc) *ValveMap{
		f:= new(ValveMap)
		f.MapFunc = mapFunc
		return f
}