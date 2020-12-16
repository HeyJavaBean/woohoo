package functions


type FilterFunc func(interface{}) bool



type ValveFilter struct{
	FilterFunc FilterFunc
}

func (valve *ValveFilter) Fire(in interface{}) []interface{}{

	out := valve.FilterFunc(in)
	if out{
		return []interface{}{in}
	}
	return nil

}

func NewFilter(filterFunc FilterFunc) *ValveFilter{
	f:= new(ValveFilter)
	f.FilterFunc = filterFunc
	return f
}