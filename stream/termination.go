package stream

import "woohoo/functions"






type ReduceFunc func(sum,next interface{}) interface{}














func (stream *Stream) Reduce(reduceFunc ReduceFunc) interface{} {

	stream.DoFireUp()

	out := stream.output

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


func (stream *Stream) GroupBy(identifyFunc IdentifyFunc) map[string][]interface{} {

	stream.DoFireUp()

	out := stream.output

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

func (stream *Stream) GroupCount(identifyFunc IdentifyFunc) map[string]int {

	stream.DoFireUp()

	out := stream.output

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
func (stream *Stream) Execute() []interface{} {

	stream.DoFireUp()

	output := []interface{}{}

	out := stream.output

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
func (stream *Stream) ToArray() []interface{} {

	stream.DoFireUp()

	output := []interface{}{}

	out := stream.output

	for {
		if data, ok := <-*out; ok {
			output = append(output, data)
		} else {
			break
		}
	}

	return output

}

func (stream *Stream) ForEach(peekFunc functions.PeekFunc){

	stream.DoFireUp()

	out := stream.output

	for {
		if data, ok := <-*out; ok {
			peekFunc(data)
		} else {
			break
		}
	}
}

func (stream *Stream) Count() int{

	stream.DoFireUp()

	out := stream.output

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


