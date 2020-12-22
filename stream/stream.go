package stream

import "github.com/HeyJavaBean/woohoo/stage"

type Stream struct {
	//数据源
	input *chan interface{}
	//处理好了的数据
	Output *chan interface{}
	//函数链头部
	Head *stage.Stage
	//函数链尾部
	Tail *stage.Stage
}

//获取一个流
func GetStream(ar []interface{}) *Stream {
	in := make(chan interface{}, len(ar))
	out := make(chan interface{}, len(ar))
	for _, a := range ar {
		in <- a
	}
	close(in)
	s := &Stream{&in, &out, nil, nil}
	return s.AddStage(true, stage.NewPassthrough())
}



func (s *Stream) AddStage(stateless bool, fun stage.StageFunc) *Stream{
	input := s.Output
	c := make(chan interface{})
	s.Output = &c
	st := &stage.Stage{input,&c,fun,nil,stateless}
	if s.Head==nil{
		st.Input = s.input
		//st.Output = s.Output
		s.Tail= st
		s.Head= st
	}else{
		s.Tail.Next = st
		s.Tail = st
	}
	return s
}


//把所有内容执行启动并且输出到输出管道里
func (s *Stream) DoFireUp() {
	stage := s.Head
	for stage != nil {
		//挨个激活
		stage.FireValve()
		stage = stage.Next
	}
}


