package main

import (
	"fmt"
)



func main() {

	arr := []interface{}{}

	for i:=0;i<8;i++{
		arr = append(arr, i)
	}

	var mappw = func(in interface{}) interface{}{
		intt := in.(int)
		intt=intt + 0
		return intt
	}
	//
	//var filterF = func(in interface{}) bool{
	//	intt := in.(int)
	//	return intt%2==0
	//}

	var flatF = func(in interface{}) []interface{}{
		intt := in.(int)
		return []interface{}{intt,intt,intt}
	}

	//start:= time.Now()
	output := GetStream(arr, len(arr)).Map(mappw).FlatMap(flatF).Execute()
	//end := time.Now()

	for _, o := range output {
		it := o.(int)
		fmt.Println(it)
	}

	//fmt.Println(end.Sub(start).Seconds())

}
