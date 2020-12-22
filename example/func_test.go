package example

import (
	"fmt"
	"github.com/HeyJavaBean/woohoo/stream"
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


func assertEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}

func TestFilter(t *testing.T) {

	lenUnder5 := func(in interface{}) bool{
		str := in.(string)
		return len(str)<5
	}

	res := stream.GetStream(getSong2lyric()).Filter(lenUnder5).ToArray()

	fmt.Println(res)
}