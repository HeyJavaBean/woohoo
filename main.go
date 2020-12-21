package main

import (
	"fmt"
	"github.com/HeyJavaBean/woohoo/stream"
	"math/rand"
)


//var mappw = func(in interface{}){
//	intt := in.(int)
//	fmt.Println(intt,":peek")
//}


var mappf = func(in interface{}) interface{}{
	intt := in.(int)
	return intt+rand.Intn(10)
}

//
//var filterF = func(in interface{}) bool{
//	intt := in.(int)
//	fmt.Println(intt,":filter")
//	return intt%2==0
//}

//var flatF = func(in interface{}) []interface{}{
//	intt := in.(int)
//	fmt.Println(intt,":flat")
//	return []interface{}{intt+2,intt-3}
//}

var comp =  func(a interface{}, b interface{}) bool{
	ai :=a.(int)
	bi :=b.(int)
	return ai<bi
}

var gb = func(in interface{}) string{
	intt := in.(int)
	if intt>20{
		return "bigger than 20"
	}else{
		return "smaller than 21"
	}
}

var rd = func(sum,next interface{}) interface{}{
	ni := next.(int)
	si := sum.(int)
	si+=ni
	return si
}

func main() {

	arr := []interface{}{}

	for i:=0;i<25;i++{
		arr = append(arr, i)
	}

	by := stream.GetStream(arr).Map(mappf).Sort(comp).Limit(3).Reduce(rd)

	fmt.Println(by)

}
