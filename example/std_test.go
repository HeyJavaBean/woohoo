package example

import (
	"fmt"
	"github.com/HeyJavaBean/woohoo/stage"
	"github.com/HeyJavaBean/woohoo/stream"
	"strings"
	"testing"
)

func TestIt(t *testing.T) {
	//对Song 2的歌词进行WordCount操作，统计除了"Woo-Hoo"以外的词频



	lyric := song2lyric()
	o := stream.GetStream(lyric).FlatMap(strToWord).Map(removeComma).ThenGroupCount(groupByWord).Max(comp)

	fmt.Println(o)


}

var comp = func(a,b interface{}) bool{
	return a.(stage.EntryCount).V>b.(stage.EntryCount).V
}

//把每句歌词转换为单词
var contentToNum = func(in interface{}) interface{}{
	e := in.(stage.Entry)
	return len(e.Vs)
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
