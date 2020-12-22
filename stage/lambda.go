package stage

type ReduceFunc func(sum,next interface{}) interface{}


type FilterFunc func(interface{}) bool

type FlatMapFunc func(interface{}) []interface{}
type MapFunc func(interface{}) interface{}
type PeekFunc func(interface{})

//感觉这个有点问题
type Comparator func(a,b interface{}) bool

