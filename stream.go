package main

import (
	"sync"
	"woohoo/functions"
)


type Stream struct {
	//数据源，全部堆积在这里
	input chan interface{}
	//处理好了的数据都放到这里来
	output chan interface{}
	//结束信号
	wg sync.WaitGroup
	//并发数量
	paraNum int
	//函数模型链
	funcChain *functions.Function
	//函数链尾部
	funcTail *functions.Function
	//开始准备
	startFlag sync.WaitGroup

}

func GetStream(ar []interface{}, paraNum int) *Stream {

	if paraNum < 1 {
		//非法输入默认设置为1
		paraNum = 1
	}

	input := make(chan interface{}, len(ar))
	//可能要考虑一下chan 长度的问题
	output := make(chan interface{}, len(ar))
	wg := sync.WaitGroup{}
	sg := sync.WaitGroup{}
	wg.Add(paraNum)
	sg.Add(1)
	for _, a := range ar {
		input <- a
	}
	close(input)

	funcH := new(functions.Function)
	funcT := funcH
	stream := Stream{input,output,wg, paraNum,funcH,funcT,sg}


	stream.doFireUp()

	return &stream
}

func (stream *Stream) Map(mapFunc functions.MapFunc) *Stream{
	f:=functions.AddValve(functions.NewMap(mapFunc))
	stream.funcTail.NextFunc = f
	stream.funcTail = f
	return stream
}

func (stream *Stream) Filter(filterFunc functions.FilterFunc) *Stream{
	f:=functions.AddValve(functions.NewFilter(filterFunc))
	stream.funcTail.NextFunc = f
	stream.funcTail = f
	return stream
}

func (stream *Stream) FlatMap(fmF functions.FlatMapFunc) *Stream{
	f:=functions.AddValve(functions.NewFlatMap(fmF))
	stream.funcTail.NextFunc = f
	stream.funcTail = f
	return stream
}

//把所有内容执行启动并且输出到输出管道里
func (stream *Stream) doFireUp(){

	//结束监听者
	go func() {
		stream.wg.Wait()
		close(stream.output)
	}()

	for i:=0;i<stream.paraNum;i++{
		go stream.fireUp()
	}
}

func (stream *Stream) fireUp(){

	stream.startFlag.Wait()

	for {
		if data, ok := <-stream.input; ok {
			//执行函数链操作
			function := stream.funcChain.NextFunc
			for function !=nil{
				out := function.Func.Fire(data)
				if out!=nil&&len(out)>0{
					function = function.NextFunc
					data = out[0]
				}else{
					break
				}
			}
		} else {
			stream.output<-data
			stream.wg.Done()
			break
		}
	}

}

//一个自定义的比较简单的终端方法，把数据全部都输出到另外一个[]interface里去
func (stream *Stream) Execute() []interface{}{

	//stream.doFireUp()

	stream.startFlag.Done()

	output := []interface{}{}

	out := stream.output

	for {
		if data, ok := <-out; ok {
			output = append(output, data)
		} else {
			break
		}
	}
	return output

}

