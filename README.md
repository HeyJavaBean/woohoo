# woohoo
A Java-Stream style functional programing api for golang

一个仿照Java Stream库的Go语言流式计算实现版本

目前还是个demo


# Example
```golang


func main() {

	//对Song 2的歌词进行WordCount操作，统计除了"Woo-Hoo"以外的词频
	wordCountMap := stream.GetStream(song2lyric()).
		FlatMap(strToWord).Map(removeComma).Filter(exceptWooHoo).GroupCount(groupByWord)

	for word, count := range wordCountMap {
		fmt.Println(word,": ",count)
	}

}

//把每句歌词转换为单词
var strToWord = func(in interface{}) []interface{}{
	str := in.(string)
	return toInterfaceArray(strings.Split(str," "))
}

//去除符号
var removeComma = func(in interface{}) interface{} {
	str := in.(string)
	str = strings.ReplaceAll(str,"(","")
	str = strings.ReplaceAll(str,")","")
	str = strings.ReplaceAll(str,",","")
	return str
}

//去除Woo-Hoo单词
var exceptWooHoo = func(in interface{}) bool {
	str := in.(string)
	return str!="Woo-hoo"
}

//按照单词进行分组
var groupByWord = func(in interface{}) string{
	return in.(string)
}

//以[]string的方式获取song 2 的歌词
func song2lyric() []interface{}{

	return  []interface{}{
		"Woo-hoo",
		"Woo-hoo",
		"Woo-hoo",
		"Woo-hoo",
		"I got my head checked",
		"By a jumbo jet",
		"It wasn't easy",
		"But nothing is",
		"No",
		"when I feel heavy metal",
		"(Woo-hoo) and I'm pins and I'm needles",
		"(Woo-hoo) well, I lie and I'm easy",
		"All of the time but I'm never sure why I need you",
		"Pleased to meet you",
		"I got my head done",
		"When I was young",
		"It's not my problem",
		"It's not my problem",
		"when I feel heavy metal",
		"(Woo-hoo) and I'm pins and I'm needles",
		"(Woo-hoo) well, I lie and I'm easy",
		"All of the time but I'm never sure why I need you",
		"Pleased to meet you",
		"Yeah, yeah",
		"Yeah, yeah",
		"Yeah, yeah",
		"Oh, yeah",
	}

}

func toInterfaceArray(strs []string) []interface{}{
	array := []interface{}{}
	for _,s := range strs {
		array = append(array, s)
	}
	return array
}

```
