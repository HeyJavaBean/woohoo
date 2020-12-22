package stream

import "woohoo/stage"

type ReduceFunc func(sum,next interface{}) interface{}












func (s *Stream) Reduce(reduceFunc ReduceFunc) interface{} {

	s.DoFireUp()

	out := s.Output

	var result interface{}

	if data, ok := <-*out; ok {
		result = data
	} else {
		return result
	}

	for {
		if data, ok := <-*out; ok {
			result = reduceFunc(result,data)
		} else {
			break
		}
	}

	return result

}


type IdentifyFunc func(in interface{}) string


func (s *Stream) GroupBy(identifyFunc IdentifyFunc) map[string][]interface{} {

	s.DoFireUp()

	out := s.Output

	res := map[string][]interface{}{}

	for {
		if data, ok := <-*out; ok {

			k := identifyFunc(data)
			res[k] = append(res[k], data)

		} else {
			break
		}
	}

	return res

}

func (s *Stream) GroupCount(identifyFunc IdentifyFunc) map[string]int {

	s.DoFireUp()

	out := s.Output

	res := map[string]int{}
	total := 0
	for {
		if data, ok := <-*out; ok {

			k := identifyFunc(data)
			res[k]++
			total++
		} else {
			break
		}
	}

	//额外增加一个总数
	res["totalNum"] = total

	return res

}








//一个自定义的比较简单的终端方法，把数据全部都输出到另外一个[]interface里去
func (s *Stream) Execute() []interface{} {

	s.DoFireUp()

	output := []interface{}{}

	out := s.Output

	for {
		if data, ok := <-*out; ok {
			output = append(output, data)
		} else {
			break
		}
	}

	return output

}



//一个自定义的比较简单的终端方法，把数据全部都输出到另外一个[]interface里去
func (s *Stream) ToArray() []interface{} {

	s.DoFireUp()

	output := []interface{}{}

	out := s.Output

	for {
		if data, ok := <-*out; ok {
			output = append(output, data)
		} else {
			break
		}
	}

	return output

}

func (s *Stream) ForEach(peekFunc stage.PeekFunc){

	s.DoFireUp()

	out := s.Output

	for {
		if data, ok := <-*out; ok {
			peekFunc(data)
		} else {
			break
		}
	}
}

func (s *Stream) Count() int{

	s.DoFireUp()

	out := s.Output

	c:=0

	for {
		if _, ok := <-*out; ok {
			c++
		} else {
			break
		}
	}
	return c
}


