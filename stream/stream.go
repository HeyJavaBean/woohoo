package stream

import (
	"woohoo/functions"
	"woohoo/termination"
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
	//终端
	termination termination.Termination
}

//获取一个流，默认进行并发执行
func GetStream(ar []interface{}) *Stream {

	//把数据全放到一个channel里等待
	input := make(chan interface{}, len(ar))
	//这里整一个有长度的ar，其实也无所谓的
	output := make(chan interface{}, len(ar))
	//数据放入源内
	for _, a := range ar {
		input <- a
	}
	//关闭
	close(input)
	//添加一个基础阀门
	funcH := functions.AddStateless(&input, &output, functions.NewPassthrough())
	funcT := funcH
	stream := Stream{&input, &output, funcH, funcT,nil}

	return &stream
}


func (stream *Stream) AddStatefulStage(fun functions.ValveFunc) *Stream {
	//上一节的输出作为本节的输入
	input := stream.output
	//准备一个输出管道
	c := make(chan interface{})
	stream.output = &c
	//把上一节的输出作为下一节的输入
	f := functions.AddStatefulValve(input, stream.output, fun)
	//这节函数加到尾巴上
	stream.funcTail.NextFunc = f
	//更新尾部节点
	stream.funcTail = f

	return stream
}

func (stream *Stream) AddStage(fun functions.ValveFunc) *Stream {
	//上一节的输出作为本节的输入
	input := stream.output
	//准备一个输出管道
	c := make(chan interface{})
	stream.output = &c
	//把上一节的输出作为下一节的输入
	f := functions.AddStateless(input, stream.output, fun)
	//这节函数加到尾巴上
	stream.funcTail.NextFunc = f
	//更新尾部节点
	stream.funcTail = f

	return stream
}

//把所有内容执行启动并且输出到输出管道里
func (stream *Stream) DoFireUp() {
	chain := stream.funcChain
	for chain != nil {
		//挨个激活
		chain.FireValve()
		chain = chain.NextFunc
	}
}


