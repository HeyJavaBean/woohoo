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
		intt=intt + 3
		fmt.Println("done!")
		return intt
	}

	var filterF = func(in interface{}) bool{
		intt := in.(int)
		fmt.Println("this is:",intt,"and ",intt%2==0)
		return intt%2==0
	}
	//
	//var flatF = func(in interface{}) []interface{}{
	//	intt := in.(int)
	//	return []interface{}{intt+2,intt-3}
	//}

	//start:= time.Now()
	output := GetStream(arr, 1).Map(mappw).Filter(filterF).Execute()
	//end := time.Now()

	fmt.Println("done!")
	for _, o := range output {
		it := o.(int)
		fmt.Println(it)
	}

	//fmt.Println(end.Sub(start).Seconds())

}
