package main

import (
	"woohoo/functions"
)

type Stream struct {
	//数据源，全部堆积在这里
	input *chan interface{}
	//处理好了的数据都放到这里来
	output *chan interface{}
	//函数模型链
	funcChain *functions.Function
	//函数链尾部
	funcTail *functions.Function
}

func GetStream(ar []interface{}, paraNum int) *Stream {

	if paraNum < 1 {
		//非法输入默认设置为1
		paraNum = 1
	}

	input := make(chan interface{}, len(ar))
	//可能要考虑一下chan 长度的问题
	output := make(chan interface{}, len(ar))
	for _, a := range ar {
		input <- a
	}
	close(input)

	funcH := functions.AddValve(&input, &output, functions.NewPassthrough())
	funcT := funcH
	stream := Stream{&input, &output, funcH, funcT}

	return &stream
}

func (stream *Stream) Map(mapFunc functions.MapFunc) *Stream {

	op := make(chan interface{})
	f := functions.AddValve(stream.output, &op, functions.NewMap(mapFunc))
	stream.funcTail.NextFunc = f
	stream.funcTail = f
	stream.output = f.Output
	return stream
}

func (stream *Stream) Filter(filterFunc functions.FilterFunc) *Stream {
	op := make(chan interface{})
	f := functions.AddValve(stream.output, &op, functions.NewFilter(filterFunc))
	stream.funcTail.NextFunc = f
	stream.funcTail = f
	stream.output = f.Output
	return stream
}

func (stream *Stream) FlatMap(fmF functions.FlatMapFunc) *Stream {
	op := make(chan interface{})
	f := functions.AddValve(stream.output, &op, functions.NewFlatMap(fmF))
	stream.funcTail.NextFunc = f
	stream.funcTail = f
	stream.output = f.Output
	return stream
}

//把所有内容执行启动并且输出到输出管道里
func (stream *Stream) doFireUp() {

	chain := stream.funcChain
	for chain != nil {
		chain.FireValve()
		chain = chain.NextFunc
	}

}

//一个自定义的比较简单的终端方法，把数据全部都输出到另外一个[]interface里去
func (stream *Stream) Execute() []interface{} {

	stream.doFireUp()

	output := []interface{}{}

	out := stream.funcTail.Output

	for {
		if data, ok := <-*out; ok {
			output = append(output, data)
		} else {
			break
		}
	}

	return output

}
