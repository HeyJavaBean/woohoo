package stage

type ReduceFunc func(sum,next interface{}) interface{}


type FilterFunc func(interface{}) bool

type FlatMapFunc func(interface{}) []interface{}
type MapFunc func(interface{}) interface{}
type PeekFunc func(interface{})

type IdentifyFunc func(in interface{}) string



//感觉这个有点问题  返回true就是从小到打，返回false就是从大到小
type Comparator func(a,b interface{}) bool

