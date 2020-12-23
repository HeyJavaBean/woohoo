package example

import (
	"fmt"
	"github.com/HeyJavaBean/woohoo/stage"
	"github.com/HeyJavaBean/woohoo/stream"
	"github.com/stretchr/testify/assert"
	"sort"
	"strconv"
	"strings"
	"testing"
)

//以[]string的方式获取song 2 的歌词
func getSong2lyric() []interface{}{

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



var lenUnder5 = func(in interface{}) bool{
	str := in.(string)
	return len(str)<5
}

var mapToLen = func(in interface{}) interface{}{
	str := in.(string)
	return len(str)
}

//把每句歌词转换为单词
var strToWord = func(in interface{}) []interface{}{
	str := in.(string)
	return toInterfaceArray(strings.Split(str," "))
}

var printOut = func(in interface{}) {
	fmt.Println(in)
}

var sortByStrLen = func(a,b interface{}) bool{
	return len(a.(string))>len(b.(string))
}

var sortEntry = func(a,b interface{}) bool{
	return len(a.(stage.Entry).Vs)>len(b.(stage.Entry).Vs)
}

var groupByStrLen = func(in interface{}) string{
	return strconv.Itoa(len(in.(string)))
}

func toInterfaceArray(strs []string) []interface{}{
	array := []interface{}{}
	for _,s := range strs {
		array = append(array, s)
	}
	return array
}

func TestFilter(t *testing.T) {

	res := stream.GetStream(getSong2lyric()).Filter(lenUnder5).ToArray()

	assert.Equal(t,[]interface{}{"No"},res)
}

func TestMap(t *testing.T) {

	res := stream.GetStream(getSong2lyric()).Map(mapToLen).ToArray()

	ins := []int{}
	for _,i := range res {
		ins = append(ins, i.(int))
	}
	sort.Ints(ins)

	assert.Equal(t,[]int{2, 7, 7, 7, 7, 8, 10, 10, 10, 14, 14, 14, 16, 18, 19, 19, 19, 19, 21, 23, 23, 34, 34, 38, 38, 49, 49},ins)
}

func TestFlatMap(t *testing.T) {

	res := stream.GetStream(getSong2lyric()).FlatMap(strToWord).ToArray()

	ins := []string{}
	for _,i := range res {
		ins = append(ins, i.(string))
	}
	sort.Strings(ins)

	assert.Equal(t,[]string{"(Woo-hoo)", "(Woo-hoo)", "(Woo-hoo)", "(Woo-hoo)", "All",
		"All", "But", "By", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I'm", "I'm",
		"I'm", "I'm", "I'm", "I'm", "I'm", "I'm", "It", "It's", "It's", "No",
		"Oh,", "Pleased", "Pleased", "When", "Woo-hoo", "Woo-hoo", "Woo-hoo",
		"Woo-hoo", "Yeah,", "Yeah,", "Yeah,", "a", "and", "and", "and", "and", "and",
		"and", "but", "but", "checked", "done", "easy", "easy", "easy", "feel", "feel", "got",
		"got", "head", "head", "heavy", "heavy", "is", "jet", "jumbo", "lie", "lie", "meet", "meet",
		"metal", "metal", "my", "my", "my", "my", "need", "need", "needles", "needles", "never", "never",
		"not", "not", "nothing", "of", "of", "pins", "pins", "problem", "problem", "sure",
		"sure", "the", "the", "time", "time", "to", "to", "was",
		"wasn't", "well,", "well,", "when", "when",
		"why", "why", "yeah", "yeah", "yeah", "yeah", "you", "you", "you", "you", "young"},ins)
}

func TestDistinct(t *testing.T) {

	res := stream.GetStream(getSong2lyric()).Distinct().ToArray()

	ins := []string{}
	for _,i := range res {
		ins = append(ins, i.(string))
	}
	sort.Strings(ins)

	assert.Equal(t,[]string{"(Woo-hoo) and I'm pins and I'm needles", "(Woo-hoo) well, " +
		"I lie and I'm easy", "All of the time but I'm never sure why I need you",
		"But nothing is", "By a jumbo jet", "I got my head checked", "I got my head done",
		"It wasn't easy", "It's not my problem", "No", "Oh, yeah",
		"Pleased to meet you", "When I was young", "Woo-hoo",
		"Yeah, yeah", "when I feel heavy metal"},ins)
}

func TestLimit(t *testing.T) {

	res := stream.GetStream(getSong2lyric()).Limit(3).ToArray()

	assert.Equal(t,[]interface{}{"Woo-hoo", "Woo-hoo", "Woo-hoo"},res)
}

func TestMeter(t *testing.T) {

	total := stage.NewCounter()

	stream.GetStream(getSong2lyric()).Meter(total).ToArray()

	assert.Equal(t,27,*total)
}

//todo 这俩还不太确定怎么写测试用例
func TestPeek(t *testing.T) {

	stream.GetStream(getSong2lyric()).Peek(printOut).ToArray()

}

//todo 这俩还不太确定怎么写测试用例
func TestSafePeek(t *testing.T) {

	stream.GetStream(getSong2lyric()).SafePeek(printOut).ToArray()

}

func TestSkip(t *testing.T) {

	res := stream.GetStream(getSong2lyric()).Skip(26).ToArray()

	assert.Equal(t,[]interface{}{"Oh, yeah"},res)
}

func TestSort(t *testing.T) {

	res := stream.GetStream(getSong2lyric()).Sort(sortByStrLen).ToArray()

	assert.Equal(t,[]interface{}{"No",
		"Woo-hoo", "Woo-hoo", "Woo-hoo", "Woo-hoo",
		"Oh, yeah",
		"Yeah, yeah", "Yeah, yeah", "Yeah, yeah",
		"But nothing is", "By a jumbo jet", "It wasn't easy",
		"When I was young", "I got my head done",
		"Pleased to meet you", "It's not my problem", "It's not my problem",
		"Pleased to meet you", "I got my head checked", "when I feel heavy metal",
		"when I feel heavy metal", "(Woo-hoo) well, I lie and I'm easy", "(Woo-hoo) well, I lie and I'm easy",
		"(Woo-hoo) and I'm pins and I'm needles", "(Woo-hoo) and I'm pins and I'm needles", "All of the time but I'm never sure why I need you",
		"All of the time but I'm never sure why I need you"},res)
}

//group和group count不好写测试用例