package example

import (
	"github.com/HeyJavaBean/woohoo/stream"
	"github.com/stretchr/testify/assert"
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


func TestFilter(t *testing.T) {

	res := stream.GetStream(getSong2lyric()).Filter(lenUnder5).ToArray()

	assert.Equal(t,[]interface{}{"No"},res)
}

func TestMap(t *testing.T) {

	res := stream.GetStream(getSong2lyric()).Filter(lenUnder5).Map(mapToLen).ToArray()

	assert.Equal(t,[]interface{}{2},res)
}
