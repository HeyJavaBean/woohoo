package main

import "woohoo/functions"

type Stream struct {
	//数据源，全部堆积在这里
	input chan interface{}
	//并发数量
	paraNum int
	//函数模型链

}


type FilterFunc func(interface{}) interface{}

type FlatMapFunc func(interface{}) []interface{}


func GetStream(ar []interface{}, paraNum int) *Stream {

	if paraNum < 1 {
		//非法输入默认设置为1
		paraNum = 1
	}

	input := make(chan interface{}, len(ar))
	for _, a := range ar {
		input <- a
	}
	close(input)

	stream := Stream{input, paraNum}

	return &stream
}


func (stream *Stream) Map(mapFunc functions.MapFunc) *Stream{



	return stream
}

func (stream *Stream) Filter(filterFunc FilterFunc) *Stream{



	return stream
}
