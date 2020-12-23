package stream

import "github.com/HeyJavaBean/woohoo/stage"

func (s *Stream) Reduce(reduceFunc stage.ReduceFunc) interface{} {

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
			result = reduceFunc(result, data)
		} else {
			break
		}
	}

	return result

}

func (s *Stream) GroupBy(identifyFunc stage.IdentifyFunc) map[string][]interface{} {

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

//第一个值是map，第二个是总共的个数
func (s *Stream) GroupCount(identifyFunc stage.IdentifyFunc) (map[string]int, int) {

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

	return res, total

}

//实现的不太好，是反向包装了一个新的comparator
func (s *Stream) Min(comparator stage.Comparator) interface{} {

	s.Sort(comparator).Limit(1).DoFireUp()

	out := s.Output

	if data, ok := <-*out; ok {
		return data
	} else {
		return nil
	}

}

func (s *Stream) Max(comparator stage.Comparator) interface{} {

	minComp := func(a,b interface{}) bool{
		return !comparator(a,b)
	}

	s.Sort(minComp).Limit(1).DoFireUp()

	out := s.Output

	if data, ok := <-*out; ok {
		return data
	} else {
		return nil
	}

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

func (s *Stream) ForEach(peekFunc stage.PeekFunc) {

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

//可以后期改成想办法做成并发的
func (s *Stream) AllMatch(filterFunc stage.FilterFunc) bool {

	s.DoFireUp()

	out := s.Output

	for {
		if data, ok := <-*out; ok {

			flag := filterFunc(data)
			//我不知道不关闭channel会不会引起内存泄漏gc不去回收的情况
			if !flag{
				return false
			}

		} else {
			break
		}
	}
	return true
}

func (s *Stream) AnyMatch(filterFunc stage.FilterFunc) bool {

	s.DoFireUp()

	out := s.Output

	for {
		if data, ok := <-*out; ok {

			flag := filterFunc(data)
			//我不知道不关闭channel会不会引起内存泄漏gc不去回收的情况
			if flag{
				return true
			}

		} else {
			break
		}
	}
	return false
}

func (s *Stream) NoneMatch(filterFunc stage.FilterFunc) bool {

	s.DoFireUp()

	out := s.Output

	for {
		if data, ok := <-*out; ok {

			flag := filterFunc(data)
			//我不知道不关闭channel会不会引起内存泄漏gc不去回收的情况
			if flag{
				return false
			}

		} else {
			break
		}
	}
	return true
}


